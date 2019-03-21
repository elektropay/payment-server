package swagger

type CertificateDetailsListResponse struct {
	Data []Certificate `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
