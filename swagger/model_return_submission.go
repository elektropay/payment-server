package swagger

type ReturnSubmission struct {
	Type_ string `json:"type,omitempty"`

	Id string `json:"id"`

	Version int32 `json:"version,omitempty"`

	OrganisationId string `json:"organisation_id"`

	Attributes *ReturnSubmissionAttributes `json:"attributes,omitempty"`

	Relationships *ReturnSubmissionRelationships `json:"relationships,omitempty"`
}
