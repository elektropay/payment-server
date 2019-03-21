package payment

import (
	"encoding/json"
	"net/http"

	"github.com/teivah/payment-server/swagger"
	"github.com/teivah/payment-server/utils"
	"go.uber.org/zap"
)

func HandlerPaymentIdDelete(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func HandlerPaymentIdGet(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func HandlerPaymentIdPut(w http.ResponseWriter, request *http.Request) {
	var payment swagger.PaymentUpdate
	err := decodeRequest(&payment, request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func HandlerPaymentsGet(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func HandlerPaymentsPost(w http.ResponseWriter, request *http.Request) {
	var payment swagger.PaymentCreation
	err := decodeRequest(&payment, request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = insert(payment)
	if err != nil {
		utils.Logger.Error("Unable to insert payment creation",
			zap.String("paymentId", payment.Data.Id),
			zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func decodeRequest(v interface{}, request *http.Request) error {
	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&v)
	if err != nil {
		utils.Logger.Error(
			"Error while decoding request", zap.Error(err))
		return err
	}

	return nil
}
