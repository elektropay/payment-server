package swagger

type ReturnPayment struct {
	Type_ string `json:"type,omitempty"`

	Id string `json:"id"`

	Version int32 `json:"version,omitempty"`

	OrganisationId string `json:"organisation_id"`

	Attributes *ReturnPaymentAttributes `json:"attributes,omitempty"`

	Relationships *ReturnPaymentRelationships `json:"relationships,omitempty"`
}
