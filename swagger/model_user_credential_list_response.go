package swagger

type UserCredentialListResponse struct {
	Data []Credential `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
