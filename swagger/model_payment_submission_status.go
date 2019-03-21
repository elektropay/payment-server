package swagger

type PaymentSubmissionStatus string

// List of PaymentSubmissionStatus
const (
	ACCEPTED            PaymentSubmissionStatus = "accepted"
	LIMIT_CHECK_PENDING PaymentSubmissionStatus = "limit_check_pending"
	LIMIT_CHECK_FAILED  PaymentSubmissionStatus = "limit_check_failed"
	LIMIT_CHECK_PASSED  PaymentSubmissionStatus = "limit_check_passed"
	RELEASED_TO_GATEWAY PaymentSubmissionStatus = "released_to_gateway"
	QUEUED_FOR_DELIVERY PaymentSubmissionStatus = "queued_for_delivery"
	DELIVERY_CONFIRMED  PaymentSubmissionStatus = "delivery_confirmed"
	DELIVERY_FAILED     PaymentSubmissionStatus = "delivery_failed"
	SUBMITTED           PaymentSubmissionStatus = "submitted"
	VALIDATION_PENDING  PaymentSubmissionStatus = "validation_pending"
)
