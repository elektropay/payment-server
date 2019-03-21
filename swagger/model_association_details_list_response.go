package swagger

type AssociationDetailsListResponse struct {
	Data []Association `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
