package swagger

type PaymentAdmissionStatus string

// List of PaymentAdmissionStatus
const (
	CONFIRMED PaymentAdmissionStatus = "confirmed"
	FAILED    PaymentAdmissionStatus = "failed"
)
