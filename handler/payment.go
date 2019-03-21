package handler

import (
	"encoding/json"
	"net/http"

	"github.com/teivah/payment-server/swagger"
	"github.com/teivah/payment-server/utils"
	"go.uber.org/zap"
)

func PaymentIdDelete(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func PaymentIdGet(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func PaymentIdPut(w http.ResponseWriter, request *http.Request) {
	var payment swagger.PaymentUpdate
	err := decodeRequest(payment, request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func PaymentsGet(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func PaymentsPost(w http.ResponseWriter, request *http.Request) {
	var payment swagger.PaymentCreation
	err := decodeRequest(payment, request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func decodeRequest(v interface{}, request *http.Request) error {
	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&v)
	if err != nil {
		utils.Logger.Error("Error while decoding request", zap.Error(err))
		return err
	}

	return nil
}
