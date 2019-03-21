package swagger

type UserAttributes struct {
	Username string `json:"username,omitempty"`

	Email string `json:"email,omitempty"`

	RoleIds []string `json:"role_ids,omitempty"`
}
