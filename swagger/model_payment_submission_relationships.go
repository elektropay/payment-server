package swagger

type PaymentSubmissionRelationships struct {
	Payment *RelationshipLinks `json:"payment,omitempty"`

	Validations *RelationshipLinks `json:"validations,omitempty"`
}
