package swagger

type PaymentAttributesDebtorParty struct {

	// Name of debtor as given with account
	AccountName string `json:"account_name,omitempty"`

	// Debtor account number. Allows upper case and numeric characters.
	AccountNumber string `json:"account_number,omitempty"`

	AccountNumberCode *AccountNumberCode `json:"account_number_code,omitempty"`

	AccountWith *BeneficiaryDebtorAccountHoldingEntity `json:"account_with,omitempty"`

	// Debtor address
	Address []string `json:"address,omitempty"`

	// Debtor birth date. Formatted according to ISO 8601 format: YYYY-MM-DD
	BirthDate string `json:"birth_date,omitempty"`

	// Debtor birth city
	BirthCity string `json:"birth_city,omitempty"`

	// Debtor birth country. ISO 3166 format country code
	BirthCountry string `json:"birth_country,omitempty"`

	// Debtor birth province
	BirthProvince string `json:"birth_province,omitempty"`

	// Country of debtor address. ISO 3166 format country code\"
	Country string `json:"country,omitempty"`

	// SWIFT BIC for ordering customer, either BIC8 or BIC11
	CustomerId string `json:"customer_id,omitempty"`

	// Code for `customer_id`
	CustomerIdCode string `json:"customer_id_code,omitempty"`

	// Debtor name
	Name string `json:"name,omitempty"`

	// Organisation identification of a debtor, in the case that the debtor is an organisation and not a private person
	OrganisationIdentification string `json:"organisation_identification,omitempty"`

	// The code that specifies the type of `organisation_identification`
	OrganisationIdentificationCode string `json:"organisation_identification_code,omitempty"`

	// Issuer of the `organisation_identification`
	OrganisationIdentificationIssuer string `json:"organisation_identification_issuer,omitempty"`
}
