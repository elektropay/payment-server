package swagger

import (
	"time"
)

type PaymentSubmissionAttributes struct {
	Status *PaymentSubmissionStatus `json:"status,omitempty"`

	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	StatusReason string `json:"status_reason,omitempty"`

	SubmissionDatetime time.Time `json:"submission_datetime,omitempty"`

	SettlementDate string `json:"settlement_date,omitempty"`

	SettlementCycle int32 `json:"settlement_cycle,omitempty"`

	RedirectedBankId string `json:"redirected_bank_id,omitempty"`

	RedirectedAccountNumber string `json:"redirected_account_number,omitempty"`

	LimitBreachStartDatetime time.Time `json:"limit_breach_start_datetime,omitempty"`

	LimitBreachEndDatetime time.Time `json:"limit_breach_end_datetime,omitempty"`

	TransactionStartDatetime time.Time `json:"transaction_start_datetime,omitempty"`
}
