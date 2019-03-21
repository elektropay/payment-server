package swagger

type PaymentAdmissionDetailsResponse struct {
	Data *PaymentAdmission `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
