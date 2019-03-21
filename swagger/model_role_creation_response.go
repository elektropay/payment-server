package swagger

type RoleCreationResponse struct {
	Data *Role `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
