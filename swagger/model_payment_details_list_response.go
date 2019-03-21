package swagger

type PaymentDetailsListResponse struct {
	Data []Payment `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
