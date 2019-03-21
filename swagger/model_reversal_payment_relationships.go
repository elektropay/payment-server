package swagger

type ReversalPaymentRelationships struct {
	Payment *RelationshipLinks `json:"payment,omitempty"`

	ReversalAdmission *RelationshipLinks `json:"reversal_admission,omitempty"`
}
