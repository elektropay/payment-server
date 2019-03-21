package swagger

type CredentialCreationResponse struct {
	Data *CredentialSecret `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
