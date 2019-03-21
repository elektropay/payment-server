package swagger

type ValidationSource string

// List of ValidationSource
const (
	PAYMENT_API       ValidationSource = "payment_api"
	PAYPORT_INTERFACE ValidationSource = "payport_interface"
	STARLING_GATEWAY  ValidationSource = "starling_gateway"
	BACS_GATEWAY      ValidationSource = "bacs_gateway"
)
