package swagger

type PayportAssociationAttributes struct {
	ParticipantId string `json:"participant_id,omitempty"`

	ParticipantType *PayportParticipantType `json:"participant_type,omitempty"`

	CustomerSendingFpsInstitution string `json:"customer_sending_fps_institution,omitempty"`

	SponsorBankId string `json:"sponsor_bank_id,omitempty"`

	SponsorAccountNumber string `json:"sponsor_account_number,omitempty"`
}
