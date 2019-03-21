package swagger

type Role struct {
	Type_ string `json:"type,omitempty"`

	Id string `json:"id,omitempty"`

	Version int32 `json:"version,omitempty"`

	OrganisationId string `json:"organisation_id,omitempty"`

	Attributes *RoleAttributes `json:"attributes,omitempty"`
}
