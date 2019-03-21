package swagger

type ReturnAdmissionDetailsResponse struct {
	Data *ReturnAdmission `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
