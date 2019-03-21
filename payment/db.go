package payment

import (
	"github.com/spf13/viper"
	"github.com/teivah/payment-server/utils"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
	"time"
)

var mongoCollection *mgo.Collection

func init() {
	// TODO Handle wrong parameter
	connectionTimeout := time.Duration(viper.GetInt("mongo.connection.timeout.ms")) * time.Millisecond
	mongoRequestTimeout := time.Duration(viper.GetInt("mongo.request.timeout.ms")) * time.Millisecond
	uri := viper.GetString("mongo.uri")
	databaseName := viper.GetString("mongo.payment.db")
	collectionName := viper.GetString("mongo.payment.collection")

	conn, err := mgo.DialWithTimeout(uri, connectionTimeout)
	if err != nil {
		utils.Logger.Fatal("Error while connecting to Mongo URI",
			zap.String("uri", uri),
			zap.Error(err))
		panic(err)
	}
	// TODO Check timeout
	conn.SetSocketTimeout(mongoRequestTimeout)
	mongoCollection = conn.DB(databaseName).C(collectionName)
}
