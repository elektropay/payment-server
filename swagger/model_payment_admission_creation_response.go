package swagger

type PaymentAdmissionCreationResponse struct {
	Data *PaymentAdmission `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
