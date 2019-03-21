package payment

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/teivah/payment-server/swagger"
	"github.com/teivah/payment-server/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	bsonx "gopkg.in/mgo.v2/bson"
	"net/http"
)

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
	var payment swagger.PaymentUpdate
	err := decodeRequest(&payment, request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), mongoRequestTimeout)
	id := payment.Data.Id
	fmt.Printf("%v\n", id)
	//hex := bson.ObjectIdHex(id)

	res, err := mongoCollection.UpdateOne(ctx,
		bson.M{"_id": bsonx.ObjectIdHex(id)},
		bson.M{"$set": bson.M(bson.M{"payload": payment.Data})},
	)

	if err != nil {
		utils.Logger.Error(logPostError,
			zap.String("paymentId", payment.Data.Id),
			zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(responsePostError)
		return
	}

	fmt.Printf("%v\n", res)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func HandlerPaymentsGet(w http.ResponseWriter, request *http.Request) {
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

	ctx, _ := context.WithTimeout(context.Background(), mongoRequestTimeout)

	res, err := mongoCollection.InsertOne(ctx,
		bson.M{"payload": payment.Data},
	)
	//	&document{
	//	Payload: payment.Data,
	//})
	if err != nil {
		utils.Logger.Error(logPostError,
			zap.String("paymentId", payment.Data.Id),
			zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(responsePostError)
		return
	}
	fmt.Printf("%v\n", res.InsertedID)

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
