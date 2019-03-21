package swagger

type PaymentSubmissionValidation struct {
	Type_ string `json:"type,omitempty"`

	Id string `json:"id"`

	Version int32 `json:"version,omitempty"`

	OrganisationId string `json:"organisation_id"`

	Attributes *PaymentSubmissionValidationAttributes `json:"attributes,omitempty"`
}
