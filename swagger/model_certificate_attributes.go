package swagger

type CertificateAttributes struct {
	Certificate string `json:"certificate"`

	IssuingCertificates []string `json:"issuing_certificates,omitempty"`
}
