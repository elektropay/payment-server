package swagger

type SubscriptionDetailsListResponse struct {
	Data []Subscription `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
