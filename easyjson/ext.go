package openrtb

type Ext struct {
	// BidRequest
	RequestAttributes []int64 `json:"ra,omitempty"`

	// User
	Consent string `json:"consent,omitempty"`

	// Regs
	GDPR int `json:"gdpr,omitempty"`
}
