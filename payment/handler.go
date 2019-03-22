package payment

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/teivah/payment-server/swagger"
	"github.com/teivah/payment-server/utils"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
)

type Envelope struct {
	Id      bson.ObjectId    `json:"id"        bson:"_id,omitempty"`
	Payload *swagger.Payment `json:"title"`
}

const (
	errorCodeBadRequest = "bad_request"
	errorCodeDb         = "error_db"
	errorMessageDb      = "Unable to persist user request."
)

func HandlerPaymentIdDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func HandlerPaymentIdGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
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
			Payload: payment.Data,
		})
	if err != nil {
		utils.Logger.Error("Unable to update payment in Mongo.",
			zap.String("paymentId", id),
			zap.Error(err))
		formatApiErrorResponse(w, http.StatusInternalServerError, &swagger.ApiError{
			ErrorCode:    errorCodeDb,
			ErrorMessage: errorMessageDb,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func HandlerPaymentsGet(w http.ResponseWriter, r *http.Request) {
	var payments []Envelope
	mongoClient.Find(nil).All(&payments)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
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
		Payload: payment.Data,
	})
	if err != nil {
		utils.Logger.Error("Unable to create payment in Mongo.", zap.Error(err))
		formatApiErrorResponse(w, http.StatusInternalServerError, &swagger.ApiError{
			ErrorCode:    errorCodeDb,
			ErrorMessage: errorMessageDb,
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

		formatApiErrorResponse(w, http.StatusBadRequest, &swagger.ApiError{
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
		utils.Logger.Error("Unable to format api error", zap.Error(err))
		return
	}

	io.WriteString(w, string(b))
}

func formatApiErrorResponse(w http.ResponseWriter, statusCode int, apiError *swagger.ApiError) {
	w.WriteHeader(statusCode)

	b, err := json.Marshal(apiError)
	if err != nil {
		utils.Logger.Error("Unable to format api error", zap.Error(err))

		// In case of an unexpected error, we try to write the error message anyway
		io.WriteString(w, apiError.ErrorMessage)
		return
	}

	io.WriteString(w, string(b))
}
