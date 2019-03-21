package swagger

type SendersCorrespondentAccountHoldingEntity struct {

	// Sender's correspondent's address
	BankAddress []string `json:"bank_address,omitempty"`

	// SWIFT BIC for sender's correspondent
	BankId string `json:"bank_id,omitempty"`

	BankIdCode *BankIdCode `json:"bank_id_code,omitempty"`

	// Sender's correspondent's name
	BankName string `json:"bank_name,omitempty"`

	// Identifier of the financial institution which services the account
	BankPartyId string `json:"bank_party_id,omitempty"`
}
