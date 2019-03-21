package swagger

type NewReturnAdmission struct {
	Type_ string `json:"type,omitempty"`

	Id string `json:"id"`

	Version int32 `json:"version,omitempty"`

	OrganisationId string `json:"organisation_id"`

	Attributes *NewReturnAdmissionAttributes `json:"attributes,omitempty"`

	Relationships *NewPaymentSubmissionRelationships `json:"relationships,omitempty"`
}
