package swagger

type PaymentAdmission struct {
	Type_ string `json:"type,omitempty"`

	Id string `json:"id"`

	Version int32 `json:"version,omitempty"`

	OrganisationId string `json:"organisation_id"`

	Attributes *PaymentAdmissionAttributes `json:"attributes,omitempty"`

	Relationships *PaymentAdmissionRelationships `json:"relationships,omitempty"`
}
