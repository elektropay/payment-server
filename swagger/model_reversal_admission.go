package swagger

type ReversalAdmission struct {
	Type_ string `json:"type,omitempty"`

	Id string `json:"id"`

	Version int32 `json:"version,omitempty"`

	OrganisationId string `json:"organisation_id"`

	Attributes *ReversalAdmissionAttributes `json:"attributes,omitempty"`

	Relationships *ReversalAdmissionRelationships `json:"relationships,omitempty"`
}
