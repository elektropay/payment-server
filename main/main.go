package main

import (
	"fmt"
	"net/http"

	"github.com/teivah/payment-server/payment"

	"github.com/spf13/viper"
	"github.com/teivah/payment-server/utils"
	"go.uber.org/zap"
)

func main() {
	router := payment.NewRouter()

	serverPort := viper.GetInt("server.port")
	utils.Sugar.Infof("Starting server on port %d", serverPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", serverPort), router)

	if err != nil {
		utils.Logger.Fatal("Unable to start server", zap.Error(err))
		panic(err)
	}
}
