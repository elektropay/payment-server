package swagger

type SepaInstantAssociationAttributes struct {
	BusinessUserDn string `json:"business_user_dn,omitempty"`

	TransportProfileId string `json:"transport_profile_id,omitempty"`

	Bic string `json:"bic,omitempty"`
}
