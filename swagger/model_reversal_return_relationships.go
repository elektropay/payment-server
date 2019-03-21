package swagger

type ReversalReturnRelationships struct {
	PaymentReturn *RelationshipLinks `json:"payment_return,omitempty"`

	ReturnReversalAdmission *RelationshipLinks `json:"return_reversal_admission,omitempty"`
}
