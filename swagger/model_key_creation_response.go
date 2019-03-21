package swagger

type KeyCreationResponse struct {
	Data *Key `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
