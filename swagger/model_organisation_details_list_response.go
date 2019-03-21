package swagger

type OrganisationDetailsListResponse struct {
	Data []Organisation `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
