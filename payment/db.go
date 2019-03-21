package payment

import (
	"context"
	"github.com/spf13/viper"
	"github.com/teivah/payment-server/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var mongoCollection *mongo.Collection
var mongoRequestTimeout time.Duration

type document struct {
	Id      bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Payload interface{}   `bson:"payload"`
}

func init() {
	// TODO Handle wrong parameter
	connectionTimeout := time.Duration(viper.GetInt("mongo.connection.timeout.ms")) * time.Millisecond
	mongoRequestTimeout = time.Duration(viper.GetInt("mongo.request.timeout.ms")) * time.Millisecond
	uri := viper.GetString("mongo.uri")
	databaseName := viper.GetString("mongo.payment.db")
	collectionName := viper.GetString("mongo.payment.collection")

	ctx, _ := context.WithTimeout(context.Background(), connectionTimeout)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		utils.Logger.Fatal("Error while connecting to Mongo URI",
			zap.String("uri", uri),
			zap.Error(err))
		panic(err)
	}

	mongoCollection = client.Database(databaseName).Collection(collectionName)
}
