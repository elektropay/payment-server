package swagger

type ReturnAdmission struct {
	Type_ string `json:"type,omitempty"`

	Id string `json:"id"`

	Version int32 `json:"version,omitempty"`

	OrganisationId string `json:"organisation_id"`

	Attributes *ReturnAdmissionAttributes `json:"attributes,omitempty"`

	Relationships *ReturnAdmissionRelationships `json:"relationships,omitempty"`
}
