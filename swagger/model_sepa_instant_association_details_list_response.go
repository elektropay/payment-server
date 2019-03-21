package swagger

type SepaInstantAssociationDetailsListResponse struct {
	Data []SepaInstantAssociation `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
