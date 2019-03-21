package swagger

type NewSepaInstantAssociation struct {
	Type_ string `json:"type,omitempty"`

	Id string `json:"id,omitempty"`

	Version int32 `json:"version,omitempty"`

	OrganisationId string `json:"organisation_id,omitempty"`

	Attributes *SepaInstantAssociationAttributes `json:"attributes,omitempty"`
}
