package openrtb

type Deal struct {
	ID          string   `json:"id"`                    // A unique identifier for the direct deal.
	BidFloor    float64  `json:"bidfloor,omitempty"`    // Minimum bid for this impression expressed in CPM.
	BidFloorCur string   `json:"bidfloorcur,omitempty"` // Currency specified using ISO-4217 alpha codes. This may be different from bid currency returned by bidder if this is allowed by the exchange.
	At          int      `json:"at,omitempty"`          // Optional override of the overall auction type of the bid request, where 1 = First Price, 2 = Second Price Plus, 3 = the value passed in bidfloor is the agreed upon deal price. Additional auction types can be defined by the exchange.
	Wseat       []string `json:"wseat,omitempty"`       // Whitelist of buyer seats allowed to bid on this deal. Seat IDs must be communicated between bidders and the exchange a priori. Omission implies no seat restrictions.
	Wadomain    []string `json:"wadomain,omitempty"`    // Array of advertiser domains (e.g., advertiser.com) allowed to bid on this deal. Omission implies no advertiser restrictions.
	//Ext         Extension `json:"ext,omitempty"`         // Placeholder for exchange-specific extensions to OpenRTB.
}
