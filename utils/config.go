package utils

import (
	"flag"
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func config() {
	flag.String("logging.level", "info", "Log level")
	flag.Int("server.port", 8080, "Server port")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		log.Fatalf("Unable to bind flags: %v", err)
		panic(err)
	}
}
