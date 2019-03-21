package swagger

type ReturnSubmissionDetailsResponse struct {
	Data *ReturnSubmission `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`
}
