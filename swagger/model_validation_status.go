package swagger

type ValidationStatus string

// List of ValidationStatus
const (
	FAILED       ValidationStatus = "failed"
	PASSED       ValidationStatus = "passed"
	NOT_ACCEPTED ValidationStatus = "not_accepted"
)
