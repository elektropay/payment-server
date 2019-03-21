package swagger

type ApprovalDetailsListResponse struct {
	Data []ApprovalRequest `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
