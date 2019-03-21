package swagger

type UserDetailsListResponse struct {
	Data []User `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
