package swagger

type SepaSctAssociationDetailsListResponse struct {
	Data []SepaSctAssociation `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
