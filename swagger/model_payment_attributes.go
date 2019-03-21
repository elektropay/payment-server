package swagger

import (
	"time"
)

type PaymentAttributes struct {
	// Amount of money moved between the instructing agent and instructed agent
	Amount string `json:"amount,omitempty"`

	BatchBookingIndicator string `json:"batch_booking_indicator,omitempty"`

	BatchId string `json:"batch_id,omitempty"`

	BatchType string `json:"batch_type,omitempty"`

	BeneficiaryParty *PaymentAttributesBeneficiaryParty `json:"beneficiary_party,omitempty"`

	// Category purpose in a coded form. Specifies the high level purpose of the instruction. Cannot be used at the same time as `category_purpose`.
	CategoryPurposeCoded string `json:"category_purpose_coded,omitempty"`

	// Category purpose in proprietary form. Specifies the high level purpose of the instruction. Cannot be used at the same time as `category_purpose_coded`.
	CategoryPurpose string `json:"category_purpose,omitempty"`

	ChargesInformation *ChargesInformation `json:"charges_information,omitempty"`

	// Unique identifier for organisations collecting payments
	ClearingId string `json:"clearing_id,omitempty"`

	// Currency of the transaction amount. Currency code as defined in [ISO 4217](http://www.iso.org/iso/home/standards/currency_codes.htm)
	Currency string `json:"currency,omitempty"`

	DebtorParty *PaymentAttributesDebtorParty `json:"debtor_party,omitempty"`

	// Unique identification, as assigned by the initiating party, to unambiguously identify the transaction. This identification is passed on, unchanged, throughout the entire end-to-end chain.
	EndToEndReference string `json:"end_to_end_reference,omitempty"`

	FileNumber string `json:"file_number,omitempty"`

	Fx *PaymentAttributesFx `json:"fx,omitempty"`

	// Unique identification, as assigned by the initiating party to unambigiously identify the transaction. This identification is an point-to-point reference and is passed on, unchanged, throughout the entire chain. Cannot includ leading, trailing or internal spaces.
	InstructionId string `json:"instruction_id,omitempty"`

	IntermediaryBank *IntermediaryBankAccountHoldingEntity `json:"intermediary_bank,omitempty"`

	// Numeric reference field, see scheme specific descriptions for usage
	NumericReference string `json:"numeric_reference,omitempty"`

	// Timestamp of when the payment instruction meets the set processing conditions. Format: YYYY-MM-DDThh:mm:ss:mmm+hh:mm
	PaymentAcceptanceDatetime time.Time `json:"payment_acceptance_datetime,omitempty"`

	// Unique identification, as assigned by the first instructing agent, to unambiguously identify the transaction that is passed on, unchanged, throughout the entire interbank chain.
	SchemeTransactionId string `json:"scheme_transaction_id,omitempty"`

	// The scheme-specific unique transaction ID. Populated by the scheme.
	UniqueSchemeId string `json:"unique_scheme_id,omitempty"`

	// Purpose of the payment in a proprietary form
	PaymentPurpose string `json:"payment_purpose,omitempty"`

	// Purpose of the payment in a coded form
	PaymentPurposeCoded string `json:"payment_purpose_coded,omitempty"`

	// Clearing infrastructure through which the payment instruction is to be processed. Default for given organisation ID is used if left empty. Has to be a valid scheme identifier.
	PaymentScheme string `json:"payment_scheme,omitempty"`

	PaymentType string `json:"payment_type,omitempty"`

	// Date on which the payment is to be debited from the debtor account. Formatted according to ISO 8601 format: YYYY-MM-DD.
	ProcessingDate string `json:"processing_date,omitempty"`

	// Date on which the payment is processed by the scheme. Only used if different from `processing_date`.
	SchemeProcessingDate string `json:"scheme_processing_date,omitempty"`

	ReceiversCorrespondent *ReceiversCorrespondentAccountHoldingEntity `json:"receivers_correspondent,omitempty"`

	// Payment reference for beneficiary use
	Reference string `json:"reference,omitempty"`

	// Regulatory reporting information
	RegulatoryReporting string `json:"regulatory_reporting,omitempty"`

	Reimbursement *ReimbursementAccountHoldingEntity `json:"reimbursement,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts receivable system provided by the debtor for the beneficiary.
	RemittanceInformation string `json:"remittance_information,omitempty"`

	// The scheme specific payment sub type
	SchemePaymentSubType string `json:"scheme_payment_sub_type,omitempty"`

	// The [scheme-specific payment type](#enumerations-scheme-payment-types)
	SchemePaymentType string `json:"scheme_payment_type,omitempty"`

	SendersCorrespondent *SendersCorrespondentAccountHoldingEntity `json:"senders_correspondent,omitempty"`

	StructuredReference *PaymentAttributesStructuredReference `json:"structured_reference,omitempty"`

	Swift *PaymentAttributesSwift `json:"swift,omitempty"`

	UltimateBeneficiary *UltimateEntity `json:"ultimate_beneficiary,omitempty"`

	UltimateDebtor *UltimateEntity `json:"ultimate_debtor,omitempty"`
}
