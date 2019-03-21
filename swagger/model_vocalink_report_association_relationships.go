package swagger

type VocalinkReportAssociationRelationships struct {
	BacsMemberCertificate *VocalinkReportAssociationCertificateRelationship `json:"bacs_member_certificate,omitempty"`

	BacsServiceUserCertificate *VocalinkReportAssociationCertificateRelationship `json:"bacs_service_user_certificate,omitempty"`

	FpsMemberCertificate *VocalinkReportAssociationCertificateRelationship `json:"fps_member_certificate,omitempty"`
}
