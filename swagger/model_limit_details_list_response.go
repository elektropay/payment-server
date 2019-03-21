package swagger

type LimitDetailsListResponse struct {
	Data []Limit `json:"data,omitempty"`

	Links *Links `json:"links,omitempty"`

	Meta *RecordCount `json:"meta,omitempty"`
}
