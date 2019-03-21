package swagger

type PaymentSubmittedEvent struct {
	PaymentSubmission *PaymentSubmission `json:"payment_submission,omitempty"`

	Payment *Payment `json:"payment,omitempty"`
}
