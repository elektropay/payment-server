package payment

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teivah/payment-server/swagger"
)

func TestPaymentToPaymentWithId(t *testing.T) {
	id := "001"
	uri := "localhost:80"
	var version int32 = 1
	paymentType := "type"
	relationships := &swagger.PaymentRelationships{
		PaymentAdmission: &swagger.PaymentRelationshipsPaymentAdmission{
			Data: []swagger.PaymentAdmission{
				{Id: "1"},
			},
		},
	}
	organisation := "organisation"
	attributes := &swagger.PaymentAttributes{
		Amount: "1.0",
	}

	payment := swagger.Payment{
		Version:        version,
		Type_:          paymentType,
		Relationships:  relationships,
		OrganisationId: organisation,
		Attributes:     attributes,
	}

	withId := paymentToPaymentWithId(id, uri, &payment)

	assert.Equal(t, id, withId.Id)
	assert.Equal(t, attributes, withId.Attributes)
	assert.Equal(t, organisation, withId.OrganisationId)
	assert.Equal(t, relationships, withId.Relationships)
	assert.Equal(t, paymentType, withId.Type_)
	assert.Equal(t, version, withId.Version)
	assert.Equal(t, "localhost:80"+paymentPrefix+"/"+id, withId.Links.Self)
}

func TestPaymentToPaymentWithIdWithNilPayment(t *testing.T) {
	id := "001"
	uri := "localhost:80"
	withId := paymentToPaymentWithId(id, uri, nil)
	assert.Equal(t, id, withId.Id)
	assert.Equal(t, "localhost:80"+paymentPrefix+"/"+id, withId.Links.Self)
}
