package swagger

type SettlementCycleAttributes struct {
	Gateway string `json:"gateway,omitempty"`

	SettlementCycleType string `json:"settlement_cycle_type,omitempty"`

	SettlementCycleNumber int32 `json:"settlement_cycle_number,omitempty"`
}
