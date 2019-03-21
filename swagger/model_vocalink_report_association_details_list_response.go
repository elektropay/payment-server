package swagger

type VocalinkReportAssociationDetailsListResponse struct {
	Data []VocalinkReportAssociation `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
