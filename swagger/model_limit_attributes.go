package swagger

type LimitAttributes struct {
	Gateway string `json:"gateway,omitempty"`

	Scheme string `json:"scheme,omitempty"`

	Amount string `json:"amount,omitempty"`

	SettlementCycleType *SettlementCycleType `json:"settlement_cycle_type,omitempty"`
}
