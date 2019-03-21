package swagger

type PaymentDetailsResponse struct {
	Data *Payment `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
