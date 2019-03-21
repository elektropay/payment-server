package payment

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/teivah/payment-server/utils"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		utils.Sugar.Infof("%s %s %s %s",
			r.Method, r.RequestURI, name, time.Since(start))
	})
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v1/",
		Index,
	},

	Route{
		"PaymentIdDelete",
		strings.ToUpper("Delete"),
		"/v1/payment/{id}",
		HandlerPaymentIdDelete,
	},

	Route{
		"PaymentIdGet",
		strings.ToUpper("Get"),
		"/v1/payment/{id}",
		HandlerPaymentIdGet,
	},

	Route{
		"PaymentIdPut",
		strings.ToUpper("Put"),
		"/v1/payment/{id}",
		HandlerPaymentIdPut,
	},

	Route{
		"PaymentsGet",
		strings.ToUpper("Get"),
		"/v1/payments",
		HandlerPaymentsGet,
	},

	Route{
		"PaymentsPost",
		strings.ToUpper("Post"),
		"/v1/payments",
		HandlerPaymentsPost,
	},
}
