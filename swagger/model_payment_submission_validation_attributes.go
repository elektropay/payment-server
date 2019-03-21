package swagger

type PaymentSubmissionValidationAttributes struct {
	Source *ValidationSource `json:"source,omitempty"`

	Status *ValidationStatus `json:"status,omitempty"`

	Response string `json:"response,omitempty"`
}
