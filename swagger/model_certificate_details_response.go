package swagger

type CertificateDetailsResponse struct {
	Data *Certificate `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
