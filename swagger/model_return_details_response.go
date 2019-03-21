package swagger

type ReturnDetailsResponse struct {
	Data *ReturnPayment `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
