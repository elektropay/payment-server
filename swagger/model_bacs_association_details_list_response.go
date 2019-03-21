package swagger

type BacsAssociationDetailsListResponse struct {
	Data []BacsAssociation `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
