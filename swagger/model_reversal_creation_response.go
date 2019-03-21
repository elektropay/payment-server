package swagger

type ReversalCreationResponse struct {
	Data *ReversalPayment `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
