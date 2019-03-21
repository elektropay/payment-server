package swagger

type PaymentRelationships struct {
	PaymentSubmission *PaymentRelationshipsPaymentSubmission `json:"payment_submission,omitempty"`

	PaymentReturn *PaymentRelationshipsPaymentReturn `json:"payment_return,omitempty"`

	PaymentAdmission *PaymentRelationshipsPaymentAdmission `json:"payment_admission,omitempty"`

	PaymentReversal *PaymentRelationshipsPaymentReversal `json:"payment_reversal,omitempty"`
}
