package swagger

type PaymentSubmissionDetailsListResponse struct {
	Data []PaymentSubmission `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
