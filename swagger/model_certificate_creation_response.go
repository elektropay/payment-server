package swagger

type CertificateCreationResponse struct {
	Data *Certificate `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
