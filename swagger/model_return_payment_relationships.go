package swagger

type ReturnPaymentRelationships struct {
	Payment *RelationshipLinks `json:"payment,omitempty"`

	ReturnAdmission *RelationshipLinks `json:"return_admission,omitempty"`

	ReturnSubmission *RelationshipLinks `json:"return_submission,omitempty"`
}
