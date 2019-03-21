package swagger

type PaymentAttributesSwiftHeader struct {
	// Destination SWIFT logical terminal address. Complete 12-character SWIFT destination, including BIC (x8), logical terminal code (x1) and branch code (x).
	Destination string `json:"destination,omitempty"`

	// The message type of the SWIFT payment, has to match `[A-Z]{2}[0-9]{3}`. Currently `MT103` is the only supported value
	MessageType string `json:"message_type,omitempty"`

	// SWIFT priority. Either `Normal` or `Priority`.
	Priority string `json:"priority,omitempty"`

	// The destination SWIFT BIC for SWIFT MT messages being sent by the payment client to SWIFT. Formatted as BIC8 or BIC11.
	Recipient string `json:"recipient,omitempty"`

	// The source SWIFT BIC for SWIFT MT messages being received by the payment client from SWIFT. Formatted as BIC8 or BIC11.
	Source string `json:"source,omitempty"`

	// Message User Reference (MUR) value, which can be up to 16 characters, and will be returned in the ACK
	UserReference string `json:"user_reference,omitempty"`
}
