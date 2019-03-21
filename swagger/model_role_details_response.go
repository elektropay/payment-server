package swagger

type RoleDetailsResponse struct {
	Data *Role `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
