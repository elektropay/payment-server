package swagger

type AceDetailsListResponse struct {
	Data []Ace `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
