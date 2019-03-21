package swagger

type KeyDetailsListResponse struct {
	Data []Key `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
