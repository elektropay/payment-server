package swagger

type SubscriptionAttributes struct {
	CallbackUri string `json:"callback_uri,omitempty"`

	CallbackTransport string `json:"callback_transport,omitempty"`

	UserId string `json:"user_id,omitempty"`

	RecordType string `json:"record_type,omitempty"`

	EventType string `json:"event_type,omitempty"`
}
