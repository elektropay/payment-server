package swagger

type Links struct {
	Self string `json:"self,omitempty"`

	First string `json:"first,omitempty"`

	Prev string `json:"prev,omitempty"`

	Next string `json:"next,omitempty"`

	Last string `json:"last,omitempty"`
}
