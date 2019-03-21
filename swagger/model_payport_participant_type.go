package swagger

type PayportParticipantType string

// List of PayportParticipantType
const (
	SETTLING     PayportParticipantType = "settling"
	NON_SETTLING PayportParticipantType = "non_settling"
)
