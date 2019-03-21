package swagger

type BacsAssociation struct {
	Type_ string `json:"type,omitempty"`

	Id string `json:"id,omitempty"`

	Version int32 `json:"version,omitempty"`

	OrganisationId string `json:"organisation_id,omitempty"`

	Attributes *BacsAssociationAttributes `json:"attributes,omitempty"`

	Relationships *BacsAssociationRelationships `json:"relationships,omitempty"`
}
