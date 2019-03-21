package swagger

type UserDetailsResponse struct {
	Data *User `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
