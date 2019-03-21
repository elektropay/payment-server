package swagger

type SettlementCycle struct {
	Id string `json:"id"`

	Attributes *SettlementCycleAttributes `json:"attributes"`
}
