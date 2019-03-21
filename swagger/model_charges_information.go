package swagger

type ChargesInformation struct {
	BearerCode string `json:"bearer_code,omitempty"`

	ReceiverChargesAmount string `json:"receiver_charges_amount,omitempty"`

	ReceiverChargesCurrency string `json:"receiver_charges_currency,omitempty"`

	SenderCharges []ChargesInformationSenderCharges `json:"sender_charges,omitempty"`
}
