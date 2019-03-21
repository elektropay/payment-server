package swagger

type KeyAttributes struct {
	Subject string `json:"subject,omitempty"`

	PrivateKey string `json:"private_key,omitempty"`

	PublicKey string `json:"public_key,omitempty"`

	Description string `json:"description,omitempty"`

	CertificateSigningRequest string `json:"certificate_signing_request,omitempty"`
}
