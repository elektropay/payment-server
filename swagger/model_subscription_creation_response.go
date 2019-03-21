package swagger

type SubscriptionCreationResponse struct {
	Data *Subscription `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
