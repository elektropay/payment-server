package swagger

type AssociationDetailsResponse struct {
	Data *Association `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
