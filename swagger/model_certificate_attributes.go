/*
 * Payment Server.
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type CertificateAttributes struct {
	Certificate string `json:"certificate"`

	IssuingCertificates []string `json:"issuing_certificates,omitempty"`
}
