package swagger

type OrganisationCreationResponse struct {
	Data *Organisation `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
