package swagger

import (
	"time"
)

type ReturnAdmissionAttributes struct {
	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	StatusReason string `json:"status_reason,omitempty"`

	AdmissionDatetime time.Time `json:"admission_datetime,omitempty"`

	SettlementDate string `json:"settlement_date,omitempty"`

	SettlementCycle int32 `json:"settlement_cycle,omitempty"`

	Status *ReturnAdmissionStatus `json:"status,omitempty"`
}
