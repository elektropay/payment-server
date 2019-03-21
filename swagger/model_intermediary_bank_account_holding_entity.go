package swagger

type IntermediaryBankAccountHoldingEntity struct {

	// Financial institution address
	BankAddress []string `json:"bank_address,omitempty"`

	// Financial institution identification
	BankId string `json:"bank_id,omitempty"`

	BankIdCode *BankIdCode `json:"bank_id_code,omitempty"`

	// Financial institution name
	BankName string `json:"bank_name,omitempty"`

	// Identifier of the financial institution which services the account
	BankPartyId string `json:"bank_party_id,omitempty"`
}
