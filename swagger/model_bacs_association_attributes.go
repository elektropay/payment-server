package swagger

type BacsAssociationAttributes struct {
	ServiceUserNumber string `json:"service_user_number,omitempty"`

	AccountNumber string `json:"account_number,omitempty"`

	SortingCode string `json:"sorting_code,omitempty"`

	AccountType int32 `json:"account_type,omitempty"`

	BankCode string `json:"bank_code,omitempty"`

	CentreNumber string `json:"centre_number,omitempty"`
}
