package swagger

import (
	"time"
)

type AuditEntryAttributes struct {
	RecordType string `json:"record_type,omitempty"`

	RecordId string `json:"record_id,omitempty"`

	ActionedBy string `json:"actioned_by,omitempty"`

	ActionTime time.Time `json:"action_time,omitempty"`

	Description string `json:"description,omitempty"`

	BeforeData *interface{} `json:"before_data,omitempty"`

	AfterData *interface{} `json:"after_data,omitempty"`
}
