package swagger

type KeyDetailsResponse struct {
	Data *Key `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
