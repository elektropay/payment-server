package swagger

type ReturnCreationResponse struct {
	Data *ReturnPayment `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
