package swagger

type SepaSctAssociation struct {
	Type_ string `json:"type,omitempty"`

	Id string `json:"id,omitempty"`

	Version int32 `json:"version,omitempty"`

	OrganisationId string `json:"organisation_id,omitempty"`

	Attributes *SepaSctAssociationAttributes `json:"attributes,omitempty"`
}
