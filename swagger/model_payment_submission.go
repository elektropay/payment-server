package swagger

type PaymentSubmission struct {
	Type_ string `json:"type,omitempty"`

	Id string `json:"id"`

	Version int32 `json:"version,omitempty"`

	OrganisationId string `json:"organisation_id"`

	Attributes *PaymentSubmissionAttributes `json:"attributes,omitempty"`

	Relationships *PaymentSubmissionRelationships `json:"relationships,omitempty"`
}
