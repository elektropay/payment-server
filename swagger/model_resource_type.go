package swagger

type ResourceType string

// List of ResourceType
const (
	PAYPORT_ASSOCIATIONS ResourceType = "payport_associations"
	LIMITS               ResourceType = "limits"
	BACS_ASSOCIATIONS    ResourceType = "bacs_associations"
)
