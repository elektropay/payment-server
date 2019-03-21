package swagger

type PaymentCreationResponse struct {
	Data *Payment `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
