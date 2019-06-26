package openrtb

type Pmp struct {
	PrivateAuction int    `json:"private_auction,omitempty"` // Indicator of auction eligibility to seats named in the Direct Deals object, where 0 = all bids are accepted, 1 = bids are restricted to the deals specified and the terms thereof.
	Deals          []Deal `json:"deals,omitempty"`           // Array of Deal objects that convey the specific deals applicable to this impression.
	//Ext            Extension `json:"ext,omitempty"`             // Placeholder for exchange-specific extensions to OpenRTB.
}
