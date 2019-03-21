/*
 * Payment Server.
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ReturnPaymentRelationships struct {
	Payment *RelationshipLinks `json:"payment,omitempty"`

	ReturnAdmission *RelationshipLinks `json:"return_admission,omitempty"`

	ReturnSubmission *RelationshipLinks `json:"return_submission,omitempty"`
}