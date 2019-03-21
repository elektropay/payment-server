package swagger

type ReturnSubmissionStatus string

// List of ReturnSubmissionStatus
const (
	ACCEPTED            ReturnSubmissionStatus = "accepted"
	LIMIT_CHECK_PENDING ReturnSubmissionStatus = "limit_check_pending"
	LIMIT_CHECK_FAILED  ReturnSubmissionStatus = "limit_check_failed"
	LIMIT_CHECK_PASSED  ReturnSubmissionStatus = "limit_check_passed"
	RELEASED_TO_GATEWAY ReturnSubmissionStatus = "released_to_gateway"
	DELIVERY_CONFIRMED  ReturnSubmissionStatus = "delivery_confirmed"
	DELIVERY_FAILED     ReturnSubmissionStatus = "delivery_failed"
	SUBMITTED           ReturnSubmissionStatus = "submitted"
	VALIDATION_PENDING  ReturnSubmissionStatus = "validation_pending"
)
