package swagger

type ReceiversCorrespondentAccountHoldingEntity struct {

	// Receiver's correspondent's address
	BankAddress []string `json:"bank_address,omitempty"`

	// SWIFT BIC for receiver's correspondent
	BankId string `json:"bank_id,omitempty"`

	BankIdCode *BankIdCode `json:"bank_id_code,omitempty"`

	// Receiver's correspondent's name
	BankName string `json:"bank_name,omitempty"`

	// Reciever's correspondent party identifier
	BankPartyId string `json:"bank_party_id,omitempty"`
}
