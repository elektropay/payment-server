package payment

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/teivah/payment-server/swagger"
)

var externalApiUri string

func init() {
	externalApiUri = fmt.Sprintf("%s:%d",
		viper.GetString("server.external.hostname"),
		viper.GetInt("server.external.port"))
}

func CleanCollection() {
	mongoClient.RemoveAll(nil)
}

func paymentToPaymentWithId(id, uri string, payment *swagger.Payment) *swagger.PaymentWithId {
	if payment == nil {
		return &swagger.PaymentWithId{
			Id: id,
			Links: &swagger.Links{
				Self: fmt.Sprintf("%s%s/%s", uri, paymentPrefix, id),
			},
		}
	}

	// To avoid having to copy the payment content, maybe composition would have been a better choice?
	return &swagger.PaymentWithId{
		Id:             id,
		Attributes:     payment.Attributes,
		OrganisationId: payment.OrganisationId,
		Relationships:  payment.Relationships,
		Type_:          payment.Type_,
		Version:        payment.Version,
		Links: &swagger.Links{
			Self: fmt.Sprintf("%s%s/%s", uri, paymentPrefix, id),
		},
	}
}
