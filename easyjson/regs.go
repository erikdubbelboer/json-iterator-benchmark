package openrtb

type Regs struct {
	COPPA int `json:"coppa,omitempty"` // Flag indicating if this request is subject to the COPPA regulations established by the USA FTC, where 0 = no, 1 = yes.
	Ext   Ext `json:"ext,omitempty"`   // Placeholder for exchange-specific extensions to OpenRTB.
}
