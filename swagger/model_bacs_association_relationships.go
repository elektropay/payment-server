package swagger

type BacsAssociationRelationships struct {
	InputCertificate *BacsAssociationCertificateRelationship `json:"input_certificate,omitempty"`

	OutputCertificate *BacsAssociationCertificateRelationship `json:"output_certificate,omitempty"`

	MessagingCertificate *BacsAssociationCertificateRelationship `json:"messaging_certificate,omitempty"`
}
