package swagger

type ReturnReversalAdmission struct {
	Type_ string `json:"type,omitempty"`

	Id string `json:"id"`

	Version int32 `json:"version,omitempty"`

	OrganisationId string `json:"organisation_id"`

	Attributes *ReturnReversalAdmissionAttributes `json:"attributes,omitempty"`

	Relationships *ReturnReversalAdmissionRelationships `json:"relationships,omitempty"`
}
