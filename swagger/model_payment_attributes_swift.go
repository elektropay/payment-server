package swagger

type PaymentAttributesSwift struct {

	// SWIFT service level
	BankOperationCode string `json:"bank_operation_code,omitempty"`

	Header *PaymentAttributesSwiftHeader `json:"header,omitempty"`

	// A SWIFT instruction code
	InstructionCode string `json:"instruction_code,omitempty"`

	// This field specifies additional information for the Receiver or other party specified.
	SenderReceiverInformation string `json:"sender_receiver_information,omitempty"`

	// This repetitive field specifies one or several time indication(s) related to the processing of the payment instruction.
	TimeIndication string `json:"time_indication,omitempty"`
}
