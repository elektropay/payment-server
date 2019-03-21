package swagger

type BacsAssociationDetailsResponse struct {
	Data *BacsAssociation `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
