/*
 * Payment Server.
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ValidationStatus string

// List of ValidationStatus
const (
	FAILED       ValidationStatus = "failed"
	PASSED       ValidationStatus = "passed"
	NOT_ACCEPTED ValidationStatus = "not_accepted"
)