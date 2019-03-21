package swagger

type RoleDetailsListResponse struct {
	Data []Role `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
