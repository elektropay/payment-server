package utils

import (
	"flag"
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func config() {
	// General
	flag.String("logging.level", "info", "Log level")
	flag.Int("server.port", 8080, "Server port")

	// MongoDB
	flag.String("mongo.uri", "mongodb://localhost:27017", "Mongo URI")
	flag.String("mongo.payment.db", "payment", "Mongo payment database name")
	flag.String("mongo.payment.collection", "payment", "Mongo payment collection name")
	flag.Int("mongo.connection.timeout.ms", 5000, "Mongo connection timeout in milliseconds")
	flag.Int("mongo.request.timeout.ms", 500, "Mongo request timeout in milliseconds")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		log.Fatalf("Unable to bind flags: %v", err)
		panic(err)
	}
}
