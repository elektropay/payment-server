/*
 * Payment Server.
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ApiError struct {
	ErrorMessage string `json:"error_message,omitempty"`

	ErrorCode string `json:"error_code,omitempty"`
}