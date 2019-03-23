package payment

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/teivah/payment-server/utils"
)

type route struct {
	name        string
	method      string
	pattern     string
	handlerFunc http.HandlerFunc
}

type routeSlice []route

const paymentApiVersion = "/v1/"
const paymentPrefix = paymentApiVersion + "payment"
const paymentsPrefix = paymentApiVersion + "payments"

func logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		utils.Sugar.Infof("%s %s %s %s",
			r.Method, r.RequestURI, name, time.Since(start))
	})
}

// NewRouter creates the payment router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.handlerFunc
		handler = logger(handler, route.name)

		router.
			Methods(route.method).
			Path(route.pattern).
			Name(route.name).
			Handler(handler)
	}

	return router
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Payment API v1")
}

var routes = routeSlice{
	route{
		"Index",
		"GET",
		paymentApiVersion,
		index,
	},

	route{
		"PaymentIdDelete",
		strings.ToUpper("Delete"),
		paymentPrefix + "/{id}",
		handlerPaymentIdDelete,
	},

	route{
		"PaymentIdGet",
		strings.ToUpper("Get"),
		paymentPrefix + "/{id}",
		handlerPaymentIdGet,
	},

	route{
		"PaymentIdPut",
		strings.ToUpper("Put"),
		paymentPrefix + "/{id}",
		handlerPaymentIdPut,
	},

	route{
		"PaymentsGet",
		strings.ToUpper("Get"),
		paymentsPrefix,
		handlerPaymentsGet,
	},

	route{
		"PaymentsPost",
		strings.ToUpper("Post"),
		paymentPrefix,
		handlerPaymentPost,
	},
}
