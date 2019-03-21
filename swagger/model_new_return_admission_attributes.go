package swagger

type NewReturnAdmissionAttributes struct {
	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	StatusReason string `json:"status_reason,omitempty"`

	SettlementDate string `json:"settlement_date,omitempty"`

	SettlementCycle int32 `json:"settlement_cycle,omitempty"`

	SourceGateway string `json:"source_gateway,omitempty"`
}
