package swagger

type PaymentAttributesFx struct {

	// Reference to the foreign exchange contract associated with the transaction
	ContractReference string `json:"contract_reference,omitempty"`

	// Factor used to convert an amount from the instructed currency into the transaction currency: i.e. to convert the `fx.original_amount`, expressed in the `fx.original_currency`, to `amount` specified in `currency`. Decimal value, represented as a string, maximum length 12. Must be > 0.
	ExchangeRate string `json:"exchange_rate,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as instructed by the initiating party. Decimal value. Must be > 0.
	OriginalAmount string `json:"original_amount,omitempty"`

	// Currency of `orginal_amount`. Currency code as defined in ISO 4217.
	OriginalCurrency string `json:"original_currency,omitempty"`
}
