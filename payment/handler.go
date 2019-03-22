package payment

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/teivah/payment-server/swagger"
	"github.com/teivah/payment-server/utils"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type PaymentEnvelope struct {
	Id      bson.ObjectId    `json:"id"        bson:"_id,omitempty"`
	Payload *swagger.Payment `json:"title"`
}

const (
	logPostError = "Unable to insert payment creation in Mongo."
)

var (
	responsePostError = []byte("Unable to handle payment creation.")
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
	parameters := mux.Vars(request)
	id := parameters["id"]

	var payment swagger.PaymentUpdate
	err := decodeRequest(&payment, request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = mongoClient.UpdateId(bson.ObjectIdHex(id),
		PaymentEnvelope{
			Payload: payment.Data,
		})
	if err != nil {
		utils.Logger.Error(logPostError,
			zap.String("paymentId", id),
			zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(responsePostError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func HandlerPaymentsGet(w http.ResponseWriter, request *http.Request) {
	var payments []PaymentEnvelope
	mongoClient.Find(nil).All(&payments)
	fmt.Printf("%v\n", payments[0].Id.Hex())

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func HandlerPaymentPost(w http.ResponseWriter, request *http.Request) {
	var payment swagger.PaymentCreation
	err := decodeRequest(&payment, request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = mongoClient.Insert(PaymentEnvelope{
		Payload: payment.Data,
	})
	if err != nil {
		utils.Logger.Error(logPostError,
			zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(responsePostError)
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
