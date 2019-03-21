package swagger

import (
	"time"
)

type ApprovalRequestAttributes struct {
	RecordType string `json:"record_type,omitempty"`

	RecordId string `json:"record_id,omitempty"`

	RecordOrgid string `json:"record_orgid,omitempty"`

	RecordVersion int32 `json:"record_version,omitempty"`

	Action string `json:"action,omitempty"`

	Status string `json:"status,omitempty"`

	ActionedBy string `json:"actioned_by,omitempty"`

	ActionTime time.Time `json:"action_time,omitempty"`

	BeforeData *interface{} `json:"before_data,omitempty"`

	AfterData *interface{} `json:"after_data,omitempty"`
}
