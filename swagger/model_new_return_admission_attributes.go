/*
 * Payment Server.
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type NewReturnAdmissionAttributes struct {
	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	StatusReason string `json:"status_reason,omitempty"`

	SettlementDate string `json:"settlement_date,omitempty"`

	SettlementCycle int32 `json:"settlement_cycle,omitempty"`

	SourceGateway string `json:"source_gateway,omitempty"`
}