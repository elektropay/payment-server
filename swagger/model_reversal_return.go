package swagger

type ReversalReturn struct {
	Type_ string `json:"type,omitempty"`

	Id string `json:"id"`

	Version int32 `json:"version,omitempty"`

	OrganisationId string `json:"organisation_id"`

	Attributes *interface{} `json:"attributes,omitempty"`

	Relationships *ReversalReturnRelationships `json:"relationships,omitempty"`
}
