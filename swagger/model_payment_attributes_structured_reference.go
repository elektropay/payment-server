package swagger

type PaymentAttributesStructuredReference struct {

	// Issuer of remittance reference
	Issuer string `json:"issuer,omitempty"`

	// Unique reference to unambiguously refer to the payment originated by the creditor, this reference enables reconciliation by the creditor upon receipt of the amount of money.
	Reference string `json:"reference,omitempty"`
}
