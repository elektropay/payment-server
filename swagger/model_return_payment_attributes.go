package swagger

import (
	"time"
)

type ReturnPaymentAttributes struct {
	Amount string `json:"amount,omitempty"`

	Currency string `json:"currency,omitempty"`

	ReturnCode string `json:"return_code,omitempty"`

	SchemeTransactionId string `json:"scheme_transaction_id,omitempty"`

	LimitBreachStartDatetime time.Time `json:"limit_breach_start_datetime,omitempty"`

	LimitBreachEndDatetime time.Time `json:"limit_breach_end_datetime,omitempty"`
}
