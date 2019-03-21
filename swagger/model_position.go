package swagger

type Position struct {
	Type_ string `json:"type,omitempty"`

	Id string `json:"id"`

	Version int32 `json:"version,omitempty"`

	OrganisationId string `json:"organisation_id"`

	Attributes *PositionAttributes `json:"attributes"`

	Links *Self `json:"links,omitempty"`
}
