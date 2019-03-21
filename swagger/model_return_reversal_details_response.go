package swagger

type ReturnReversalDetailsResponse struct {
	Data *ReversalReturn `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
