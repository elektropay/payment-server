package main

import (
	"github.com/teivah/payment-server/router"
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")

	router := router.NewRouter()

	http.ListenAndServe(":8080", router)
}
