package swagger

type OrganisationDetailsResponse struct {
	Data *Organisation `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
