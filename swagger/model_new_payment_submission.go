package swagger

type NewPaymentSubmission struct {
	Type_ string `json:"type,omitempty"`

	Id string `json:"id"`

	Version int32 `json:"version,omitempty"`

	OrganisationId string `json:"organisation_id"`

	Relationships *NewPaymentSubmissionRelationships `json:"relationships,omitempty"`
}
