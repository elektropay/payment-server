package swagger

type ApiError struct {
	ErrorMessage string `json:"error_message,omitempty"`

	ErrorCode string `json:"error_code,omitempty"`
}
