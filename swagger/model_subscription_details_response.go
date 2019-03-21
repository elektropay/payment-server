package swagger

type SubscriptionDetailsResponse struct {
	Data *Subscription `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
