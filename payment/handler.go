package payment

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teivah/payment-server/swagger"
	"github.com/teivah/payment-server/utils"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	errorCodeBadRequest = "bad_request"
	errorCodeHandler    = "handler_error"
	errorMessageHandler = "Unable to handle user request."
)

// handlerPaymentIdDelete handles /DELETE payment
func handlerPaymentIdDelete(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	id := parameters["id"]

	hex, err := toBsonObjectId(id)
	if err != nil {
		utils.Logger.Warn(
			"Error while decoding id", zap.Error(err))

		createErrorResponse(w, http.StatusBadRequest, &swagger.ApiError{
			ErrorCode:    errorCodeBadRequest,
			ErrorMessage: err.Error(),
		})
		return
	}

	err = mongoClient.RemoveId(hex)
	if err != nil {
		if err == mgo.ErrNotFound {
			utils.Logger.Warn("Unable to delete payment in Mongo: payment not found.",
				zap.String("paymentId", id),
				zap.Error(err))
			createErrorResponse(w, http.StatusNotFound, nil)
			return
		}

		createErrorResponse(w, http.StatusInternalServerError, &swagger.ApiError{
			ErrorCode:    errorCodeHandler,
			ErrorMessage: errorMessageHandler,
		})
		return
	}

	createPaymentResponse(w, http.StatusNoContent, id, externalApiUri, nil, nil)
}

// handlerPaymentIdGet handles GET payment
func handlerPaymentIdGet(w http.ResponseWriter, r *http.Request) {
	var envelope Envelope

	parameters := mux.Vars(r)
	id := parameters["id"]

	hex, err := toBsonObjectId(id)
	if err != nil {
		utils.Logger.Warn(
			"Error while decoding id", zap.Error(err))

		createErrorResponse(w, http.StatusBadRequest, &swagger.ApiError{
			ErrorCode:    errorCodeBadRequest,
			ErrorMessage: err.Error(),
		})
		return
	}

	err = mongoClient.FindId(hex).One(&envelope)
	if err != nil {
		if err == mgo.ErrNotFound {
			utils.Logger.Warn("Unable to update payment in Mongo: payment not found.",
				zap.String("paymentId", id),
				zap.Error(err))
			createErrorResponse(w, http.StatusNotFound, nil)
			return
		}

		createErrorResponse(w, http.StatusInternalServerError, &swagger.ApiError{
			ErrorCode:    errorCodeHandler,
			ErrorMessage: errorMessageHandler,
		})
		return
	}

	createPaymentResponse(w, http.StatusOK, id, externalApiUri, envelope.Payment,
		func(paymentWithId *swagger.PaymentWithId) interface{} {
			return paymentWithId
		})
}

// handlerPaymentIdPut handles PUT payment
func handlerPaymentIdPut(w http.ResponseWriter, r *http.Request) {
	var payment swagger.PaymentUpdate
	err := decodeRequest(&payment, w, r)
	if err != nil {
		return
	}

	parameters := mux.Vars(r)
	id := parameters["id"]
	hex, err := toBsonObjectId(id)
	if err != nil {
		utils.Logger.Warn(
			"Error while decoding id", zap.Error(err))

		createErrorResponse(w, http.StatusBadRequest, &swagger.ApiError{
			ErrorCode:    errorCodeBadRequest,
			ErrorMessage: err.Error(),
		})
		return
	}

	err = mongoClient.UpdateId(hex,
		Envelope{
			Payment: payment.Data,
		})

	if err != nil {
		if err == mgo.ErrNotFound {
			utils.Logger.Warn("Unable to update payment in Mongo: payment not found.",
				zap.String("paymentId", id),
				zap.Error(err))
			createErrorResponse(w, http.StatusNotFound, nil)
			return
		}

		utils.Logger.Error("Unable to update payment in Mongo.",
			zap.String("paymentId", id),
			zap.Error(err))
		createErrorResponse(w, http.StatusInternalServerError, &swagger.ApiError{
			ErrorCode:    errorCodeHandler,
			ErrorMessage: errorMessageHandler,
		})
		return
	}

	createPaymentResponse(w, http.StatusCreated, id, externalApiUri, payment.Data,
		func(paymentWithId *swagger.PaymentWithId) interface{} {
			return swagger.PaymentUpdateResponse{
				Data: paymentWithId,
			}
		})
}

// handlerPaymentsGet handles GET payments
func handlerPaymentsGet(w http.ResponseWriter, r *http.Request) {
	var envelopes []Envelope

	err := mongoClient.Find(nil).All(&envelopes)
	if err != nil {
		createErrorResponse(w, http.StatusInternalServerError, &swagger.ApiError{
			ErrorCode:    errorCodeHandler,
			ErrorMessage: errorMessageHandler,
		})
	}

	createPaymentsResponse(w, http.StatusOK, externalApiUri, envelopes)
}

// handlerPaymentPost handles POST payment
func handlerPaymentPost(w http.ResponseWriter, r *http.Request) {
	var payment swagger.PaymentCreation
	err := decodeRequest(&payment, w, r)
	if err != nil {
		return
	}

	id := bson.NewObjectId()
	err = mongoClient.Insert(Envelope{
		Id:      id,
		Payment: payment.Data,
	})
	if err != nil {
		utils.Logger.Error("Unable to create payment in Mongo.", zap.Error(err))
		createErrorResponse(w, http.StatusInternalServerError, &swagger.ApiError{
			ErrorCode:    errorCodeHandler,
			ErrorMessage: errorMessageHandler,
		})
		return
	}

	createPaymentResponse(w, http.StatusCreated, id.Hex(), externalApiUri, payment.Data,
		func(paymentWithId *swagger.PaymentWithId) interface{} {
			return swagger.PaymentCreationResponse{
				Data: paymentWithId,
			}
		})
}

// decodeRequest decodes an incoming HTTP request
// In case of a problem, it writes a bad request status in the http.ResponseWriter
func decodeRequest(v interface{}, w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&v)
	if err != nil {
		utils.Logger.Warn(
			"Error while decoding request", zap.Error(err))

		createErrorResponse(w, http.StatusBadRequest, &swagger.ApiError{
			ErrorCode:    errorCodeBadRequest,
			ErrorMessage: err.Error(),
		})
		return err
	}

	return nil
}

// createPaymentResponse creates a payment response
// It takes an objectResponse to wrap the response in a higher level structure like PaymentCreationResponse or PaymentUpdateResponse
func createPaymentResponse(w http.ResponseWriter, status int, id, uri string, payment *swagger.Payment,
	objectResponse func(*swagger.PaymentWithId) interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	if objectResponse == nil {
		return
	}
	withId := paymentToPaymentWithId(id, uri, payment)
	object := objectResponse(withId)
	b, err := json.Marshal(object)
	if err != nil {
		utils.Logger.Error("Unable to format payment response", zap.Error(err))
		return
	}

	io.WriteString(w, string(b))
}

// createPaymentResponse creates a payments response
func createPaymentsResponse(w http.ResponseWriter, status int, uri string, envelopes []Envelope) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	len := len(envelopes)
	withIds := make([]swagger.PaymentWithId, len)
	for i := 0; i < len; i++ {
		envelope := envelopes[i]
		withIds[i] = *paymentToPaymentWithId(envelope.Id.Hex(), uri, envelope.Payment)
	}

	response := swagger.PaymentDetailsListResponse{
		Data: withIds,
		Links: &swagger.Links{
			Self: fmt.Sprintf("%s%s", uri, paymentsPrefix),
		},
	}

	b, err := json.Marshal(response)
	if err != nil {
		utils.Logger.Error("Unable to format payments", zap.Error(err))
		return
	}

	io.WriteString(w, string(b))
}

// createErrorResponse creates a response in case of an error
// apiError input is optional
func createErrorResponse(w http.ResponseWriter, statusCode int, apiError *swagger.ApiError) {
	w.WriteHeader(statusCode)

	if apiError == nil {
		return
	}

	b, err := json.Marshal(apiError)
	if err != nil {
		utils.Logger.Error("Unable to format api error", zap.Error(err))
		// In case of an unexpected error, we try to write the error message anyway
		io.WriteString(w, apiError.ErrorMessage)
		return
	}

	io.WriteString(w, string(b))
}

// toBsonObjectId converts in string identifier in a bson.ObjectId
func toBsonObjectId(id string) (hex bson.ObjectId, err error) {
	// bson.ObjectIdHex panics if the identifier is invalid
	// We need to recover from this panic and returns a proper error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("invalid id")
		}
	}()
	return bson.ObjectIdHex(id), nil
}
