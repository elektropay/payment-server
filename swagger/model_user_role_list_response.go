package swagger

type UserRoleListResponse struct {
	Data []Role `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
