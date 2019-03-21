package payment

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/teivah/payment-server/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"time"
)

var collection *mongo.Collection
var requestTimeout time.Duration

func init() {
	// TODO Handle wrong parameter
	connectionTimeout := time.Duration(viper.GetInt("mongo.connection.timeout.ms")) * time.Millisecond
	requestTimeout = time.Duration(viper.GetInt("mongo.request.timeout.ms")) * time.Millisecond
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

	collection = client.Database(databaseName).Collection(collectionName)
}

func insert(document interface{}) error {
	ctx, _ := context.WithTimeout(context.Background(), requestTimeout)
	res, err := collection.InsertOne(ctx, document)

	if err != nil {
		return nil
	}

	id := res.InsertedID
	fmt.Printf("%v\n", id)
	return nil
}
