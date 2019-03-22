package payment

import (
	"encoding/json"
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
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func HandlerPaymentIdGet(w http.ResponseWriter, r *http.Request) {
	var envelope Envelope

	parameters := mux.Vars(r)
	id := parameters["id"]

	err := mongoClient.FindId(bson.ObjectIdHex(id)).One(&envelope)
	if err != nil {
		if err == mgo.ErrNotFound {
			utils.Logger.Warn("Unable to update payment in Mongo: payment not found.",
				zap.String("paymentId", id),
				zap.Error(err))
			formatErrorResponse(w, http.StatusNotFound, nil)
		} else {
			formatErrorResponse(w, http.StatusInternalServerError, &swagger.ApiError{
				ErrorCode:    errorCodeHandler,
				ErrorMessage: errorMessageHandler,
			})
		}
		return
	}

	formatPaymentWithIdResponse(w, id, externalApiUri, envelope.Payment)
}

func HandlerPaymentIdPut(w http.ResponseWriter, r *http.Request) {
	var payment swagger.PaymentUpdate
	err := decodeRequest(&payment, w, r)
	if err != nil {
		return
	}

	parameters := mux.Vars(r)
	id := parameters["id"]

	err = mongoClient.UpdateId(bson.ObjectIdHex(id),
		Envelope{
			Payment: payment.Data,
		})

	if err != nil {
		if err == mgo.ErrNotFound {
			utils.Logger.Warn("Unable to update payment in Mongo: payment not found.",
				zap.String("paymentId", id),
				zap.Error(err))
			formatErrorResponse(w, http.StatusNotFound, nil)
		} else {
			utils.Logger.Error("Unable to update payment in Mongo.",
				zap.String("paymentId", id),
				zap.Error(err))
			formatErrorResponse(w, http.StatusInternalServerError, &swagger.ApiError{
				ErrorCode:    errorCodeHandler,
				ErrorMessage: errorMessageHandler,
			})
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
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

	formatPaymentsWithIdResponse(w, externalApiUri, envelopes)
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

	formatPaymentWithIdResponse(w, id.Hex(), externalApiUri, payment.Data)
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

func formatPaymentWithIdResponse(w http.ResponseWriter, id, uri string, payment *swagger.Payment) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	withId := paymentToPaymentWithId(id, uri, payment)

	b, err := json.Marshal(withId)
	if err != nil {
		utils.Logger.Error("Unable to format payment response", zap.Error(err))
		return
	}

	io.WriteString(w, string(b))
}

func formatPaymentsWithIdResponse(w http.ResponseWriter, uri string, envelopes []Envelope) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

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
