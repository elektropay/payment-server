package swagger

type Payment struct {

	// Name of the resource type
	Type_ string `json:"type,omitempty"`

	// Unique resource ID
	Id string `json:"id"`

	// Version number
	Version int32 `json:"version,omitempty"`

	// Unique ID of the organisation this resource is created by
	OrganisationId string `json:"organisation_id"`

	Attributes *PaymentAttributes `json:"attributes"`

	Relationships *PaymentRelationships `json:"relationships,omitempty"`
}
