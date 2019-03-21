package swagger

type ReimbursementAccountHoldingEntity struct {

	// Third party reimbursement institution address
	BankAddress []string `json:"bank_address,omitempty"`

	// Identification of third party reimbursement institution
	BankId string `json:"bank_id,omitempty"`

	BankIdCode *BankIdCode `json:"bank_id_code,omitempty"`

	// Third party reimbursement institution name
	BankName string `json:"bank_name,omitempty"`

	// Third party reimbursement institution identifier
	BankPartyId string `json:"bank_party_id,omitempty"`
}
