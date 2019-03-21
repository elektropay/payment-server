package swagger

type AceDetailsResponse struct {
	Data *Ace `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
