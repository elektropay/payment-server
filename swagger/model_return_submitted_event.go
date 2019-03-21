package swagger

type ReturnSubmittedEvent struct {
	Payment *Payment `json:"payment,omitempty"`

	ReturnPayment *ReturnPayment `json:"return_payment,omitempty"`

	ReturnSubmission *ReturnSubmission `json:"return_submission,omitempty"`
}
