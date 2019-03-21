package swagger

type RelationshipsResponse struct {
	Data []Relationship `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
