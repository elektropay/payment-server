package swagger

type CredentialSecret struct {
	ClientId string `json:"client_id,omitempty"`

	ClientSecret string `json:"client_secret,omitempty"`
}
