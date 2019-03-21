package swagger

type AceAttributes struct {
	RoleId string `json:"role_id,omitempty"`

	Action string `json:"action,omitempty"`

	RecordType string `json:"record_type,omitempty"`

	Filter string `json:"filter,omitempty"`
}
