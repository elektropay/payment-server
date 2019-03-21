package swagger

type UltimateEntity struct {
	Name string `json:"name,omitempty"`

	Country string `json:"country,omitempty"`

	Address string `json:"address,omitempty"`

	OrganisationIdentification string `json:"organisation_identification,omitempty"`

	OrganisationIdentificationCode string `json:"organisation_identification_code,omitempty"`

	OrganisationIdentificationIssuer string `json:"organisation_identification_issuer,omitempty"`

	BirthDate string `json:"birth_date,omitempty"`

	BirthCity string `json:"birth_city,omitempty"`

	BirthCountry string `json:"birth_country,omitempty"`

	BirthProvince string `json:"birth_province,omitempty"`
}
