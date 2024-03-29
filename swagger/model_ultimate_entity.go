/*
 * Payment Server.
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

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
