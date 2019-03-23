package payment

import (
	"github.com/teivah/payment-server/swagger"
	"gopkg.in/mgo.v2/bson"
	"time"

	"github.com/spf13/viper"
	"github.com/teivah/payment-server/utils"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
)

var mongoClient *mgo.Collection

// Envelope is the representation of a Mongo document
type Envelope struct {
	// Id is the identifier
	Id bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	// Payment is the payload
	Payment *swagger.Payment `json:"payment"`
}

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
	mongoClient = conn.DB(databaseName).C(collectionName)
}
