package payment

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/teivah/payment-server/swagger"
	"github.com/teivah/payment-server/utils"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
)

type Envelope struct {
	Id      bson.ObjectId    `json:"id"        bson:"_id,omitempty"`
	Payment *swagger.Payment `json:"payment"`
}

const (
	errorCodeBadRequest = "bad_request"
	errorCodeHandler    = "handler_error"
	errorMessageHandler = "Unable to handle user request."
)

func HandlerPaymentIdDelete(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	id := parameters["id"]

	hex, err := mapIdToHex(id)
	if err != nil {
		utils.Logger.Warn(
			"Error while decoding id", zap.Error(err))

		formatErrorResponse(w, http.StatusBadRequest, &swagger.ApiError{
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
			formatErrorResponse(w, http.StatusNotFound, nil)
			return
		}

		formatErrorResponse(w, http.StatusInternalServerError, &swagger.ApiError{
			ErrorCode:    errorCodeHandler,
			ErrorMessage: errorMessageHandler,
		})
		return
	}

	formatPaymentWithIdResponse(w, http.StatusNoContent, id, externalApiUri, nil, nil)
}

func HandlerPaymentIdGet(w http.ResponseWriter, r *http.Request) {
	var envelope Envelope

	parameters := mux.Vars(r)
	id := parameters["id"]

	hex, err := mapIdToHex(id)
	if err != nil {
		utils.Logger.Warn(
			"Error while decoding id", zap.Error(err))

		formatErrorResponse(w, http.StatusBadRequest, &swagger.ApiError{
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
			formatErrorResponse(w, http.StatusNotFound, nil)
			return
		}

		formatErrorResponse(w, http.StatusInternalServerError, &swagger.ApiError{
			ErrorCode:    errorCodeHandler,
			ErrorMessage: errorMessageHandler,
		})
		return
	}

	formatPaymentWithIdResponse(w, http.StatusOK, id, externalApiUri, envelope.Payment,
		func(paymentWithId *swagger.PaymentWithId) interface{} {
			return paymentWithId
		})
}

func HandlerPaymentIdPut(w http.ResponseWriter, r *http.Request) {
	var payment swagger.PaymentUpdate
	err := decodeRequest(&payment, w, r)
	if err != nil {
		return
	}

	parameters := mux.Vars(r)
	id := parameters["id"]
	hex, err := mapIdToHex(id)
	if err != nil {
		utils.Logger.Warn(
			"Error while decoding id", zap.Error(err))

		formatErrorResponse(w, http.StatusBadRequest, &swagger.ApiError{
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
			formatErrorResponse(w, http.StatusNotFound, nil)
		}

		utils.Logger.Error("Unable to update payment in Mongo.",
			zap.String("paymentId", id),
			zap.Error(err))
		formatErrorResponse(w, http.StatusInternalServerError, &swagger.ApiError{
			ErrorCode:    errorCodeHandler,
			ErrorMessage: errorMessageHandler,
		})
		return
	}

	formatPaymentWithIdResponse(w, http.StatusCreated, id, externalApiUri, payment.Data,
		func(paymentWithId *swagger.PaymentWithId) interface{} {
			return swagger.PaymentUpdateResponse{
				Data: paymentWithId,
			}
		})
}

func HandlerPaymentsGet(w http.ResponseWriter, r *http.Request) {
	var envelopes []Envelope

	err := mongoClient.Find(nil).All(&envelopes)
	if err != nil {
		formatErrorResponse(w, http.StatusInternalServerError, &swagger.ApiError{
			ErrorCode:    errorCodeHandler,
			ErrorMessage: errorMessageHandler,
		})
	}

	formatPaymentsWithIdResponse(w, http.StatusOK, externalApiUri, envelopes)
}

func HandlerPaymentPost(w http.ResponseWriter, r *http.Request) {
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
		formatErrorResponse(w, http.StatusInternalServerError, &swagger.ApiError{
			ErrorCode:    errorCodeHandler,
			ErrorMessage: errorMessageHandler,
		})
		return
	}

	formatPaymentWithIdResponse(w, http.StatusCreated, id.Hex(), externalApiUri, payment.Data,
		func(paymentWithId *swagger.PaymentWithId) interface{} {
			return swagger.PaymentCreationResponse{
				Data: paymentWithId,
			}
		})
}

func decodeRequest(v interface{}, w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&v)
	if err != nil {
		utils.Logger.Warn(
			"Error while decoding request", zap.Error(err))

		formatErrorResponse(w, http.StatusBadRequest, &swagger.ApiError{
			ErrorCode:    errorCodeBadRequest,
			ErrorMessage: err.Error(),
		})
		return err
	}

	return nil
}

func formatPaymentWithIdResponse(w http.ResponseWriter, status int, id, uri string, payment *swagger.Payment,
	objectResponse func(*swagger.PaymentWithId) interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	withId := paymentToPaymentWithId(id, uri, payment)
	object := objectResponse(withId)
	b, err := json.Marshal(object)
	if err != nil {
		utils.Logger.Error("Unable to format payment response", zap.Error(err))
		return
	}

	io.WriteString(w, string(b))
}

func formatPaymentsWithIdResponse(w http.ResponseWriter, status int, uri string, envelopes []Envelope) {
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

func formatErrorResponse(w http.ResponseWriter, statusCode int, apiError *swagger.ApiError) {
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

func mapIdToHex(id string) (hex bson.ObjectId, err error) {
	// bson.ObjectIdHex panics if the id is invalid
	// We need to recover from this panic and returns a proper error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("invalid id")
		}
	}()
	return bson.ObjectIdHex(id), nil
}
