package swagger

type ReversalDetailsResponse struct {
	Data *ReversalPayment `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
