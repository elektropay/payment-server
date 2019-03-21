package swagger

import (
	"time"
)

type PaymentAdmissionAttributes struct {
	Status *PaymentAdmissionStatus `json:"status,omitempty"`

	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	StatusReason *PaymentAdmissionStatusReason `json:"status_reason,omitempty"`

	AdmissionDatetime time.Time `json:"admission_datetime,omitempty"`

	SettlementDate string `json:"settlement_date,omitempty"`

	SettlementCycle int32 `json:"settlement_cycle,omitempty"`
}
