package swagger

type UserCreationResponse struct {
	Data *User `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
