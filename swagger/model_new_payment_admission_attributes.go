package swagger

import (
	"time"
)

type NewPaymentAdmissionAttributes struct {
	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	AdmissionDatetime time.Time `json:"admission_datetime,omitempty"`

	SettlementDate string `json:"settlement_date,omitempty"`

	SettlementCycle int32 `json:"settlement_cycle,omitempty"`

	SourceGateway string `json:"source_gateway,omitempty"`
}
