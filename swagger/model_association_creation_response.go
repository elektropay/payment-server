package swagger

type AssociationCreationResponse struct {
	Data *Association `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
