package swagger

type PaymentSubmissionCreationResponse struct {
	Data *PaymentSubmission `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
