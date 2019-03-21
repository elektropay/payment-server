package swagger

type PayportAssociationDetailsListResponse struct {
	Data []PayportAssociation `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
