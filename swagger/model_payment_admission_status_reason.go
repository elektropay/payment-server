package swagger

type PaymentAdmissionStatusReason string

// List of PaymentAdmissionStatusReason
const (
	ACCEPTED                    PaymentAdmissionStatusReason = "accepted"
	INVALID_BENEFICIARY_DETAILS PaymentAdmissionStatusReason = "invalid_beneficiary_details"
	BANKID_NOT_PROVISIONED      PaymentAdmissionStatusReason = "bankid_not_provisioned"
	UNKNOWN_ACCOUNTNUMBER       PaymentAdmissionStatusReason = "unknown_accountnumber"
)
