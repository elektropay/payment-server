package swagger

type ReturnAdmissionRelationships struct {
	PaymentReturn *RelationshipLinks `json:"payment_return,omitempty"`

	Validations *RelationshipLinks `json:"validations,omitempty"`
}
