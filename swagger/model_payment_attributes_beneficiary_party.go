package swagger

type PaymentAttributesBeneficiaryParty struct {

	// Name of beneficiary as given with account
	AccountName string `json:"account_name,omitempty"`

	// Beneficiary account number
	AccountNumber string `json:"account_number,omitempty"`

	AccountNumberCode *AccountNumberCode `json:"account_number_code,omitempty"`

	// The type of the account given with `beneficiary_party.account_number`. Single digit number. Only required if requested by the beneficiary party. Defaults to 0.
	AccountType int32 `json:"account_type,omitempty"`

	AccountWith *BeneficiaryDebtorAccountHoldingEntity `json:"account_with,omitempty"`

	// Beneficiary address
	Address []string `json:"address,omitempty"`

	// Beneficiary birth date. Formatted according to ISO 8601 format: YYYY-MM-DD
	BirthDate string `json:"birth_date,omitempty"`

	// Beneficiary birth city
	BirthCity string `json:"birth_city,omitempty"`

	// Beneficiary birth country, ISO 3166 format country code
	BirthCountry string `json:"birth_country,omitempty"`

	// Beneficiary birth province
	BirthProvince string `json:"birth_province,omitempty"`

	// Country of the beneficiary address, ISO 3166 format country code
	Country string `json:"country,omitempty"`

	// Beneficiary name
	Name string `json:"name,omitempty"`

	// Organisation identification of a beneficiary, used in the case that the beneficiary is an organisation and not a private person
	OrganisationIdentification string `json:"organisation_identification,omitempty"`

	// The code that specifies the type of `organisation_identification`
	OrganisationIdentificationCode string `json:"organisation_identification_code,omitempty"`

	// Issuer of the organisation identification
	OrganisationIdentificationIssuer string `json:"organisation_identification_issuer,omitempty"`

	// Beneficiary phone number
	TelephoneNumber string `json:"telephone_number,omitempty"`
}
