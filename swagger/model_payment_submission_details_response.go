package swagger

type PaymentSubmissionDetailsResponse struct {
	Data *PaymentSubmission `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
