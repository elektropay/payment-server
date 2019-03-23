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

// CleanCollection clears the current Mongo collection
func CleanCollection() {
	mongoClient.RemoveAll(nil)
}

// paymentToPaymentWithId converts a Payment in a PaymentWithId struture
func paymentToPaymentWithId(id, uri string, payment *swagger.Payment) *swagger.PaymentWithId {
	if payment == nil {
		return &swagger.PaymentWithId{
			Id: id,
			Links: &swagger.Links{
				Self: formatPaymentSelfLink(uri, paymentPrefix, id),
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
			Self: formatPaymentSelfLink(uri, paymentPrefix, id),
		},
	}
}

// formatPaymentSelfLink formats the payment self link used for HATEOAS responses
func formatPaymentSelfLink(uri, paymentPrefix, id string) string {
	return fmt.Sprintf("%s%s/%s", uri, paymentPrefix, id)
}
