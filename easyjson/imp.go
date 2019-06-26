package openrtb

type Imp struct {
	ID                RawString `json:"id"`                          // A unique identifier for this impression within the context of the bid request (typically, starts with 1 and increments)
	Banner            *Banner   `json:"banner,omitempty"`            // A Banner object; required if this impression is offered as a banner ad opportunity.
	Video             *Video    `json:"video,omitempty"`             // A Video object; required if this impression is offered as a video ad opportunity.
	DisplayManager    string    `json:"displaymanager,omitempty"`    // Name of ad mediation partner, SDK technology, or player responsible for rendering ad (typically video or mobile). Used by some ad servers to customize ad code by partner. Recommended for video and/or apps.
	DisplayManagerVer string    `json:"displaymanagerver,omitempty"` // Version of ad mediation partner, SDK technology, or player responsible for rendering ad (typically video or mobile). Used by some ad servers to customize ad code by partner. Recommended for video and/or apps.
	Instl             int       `json:"instl,omitempty"`             // 1 = the ad is interstitial or full screen, 0 = not interstitial.
	TagID             string    `json:"tagid,omitempty"`             // Identifier for specific ad placement or ad tag that was used to initiate the auction. This can be useful for debugging of any issues, or for optimization by the buyer.
	BidFloor          float64   `json:"bidfloor,omitempty"`          // Minimum bid for this impression expressed in CPM.
	BidFloorCur       string    `json:"bidfloorcur,omitempty"`       // Currency specified using ISO-4217 alpha codes. This may be different from bid currency returned by bidder if this is allowed by the exchange.
	Secure            int       `json:"secure,omitempty"`            // Flag to indicate if the impression requires secure HTTPS URL creative assets and markup, where 0 = non-secure, 1 = secure. If omitted, the secure state is unknown, but non-secure HTTP support can be assumed.
	IframeBuster      []string  `json:"iframebuster,omitempty"`      // Array of exchange-specific names of supported iframe busters.
	PMP               *Pmp      `json:"pmp,omitempty"`               // A Pmp object containing any private marketplace deals in effect for this impression.
	//Ext               Extension `json:"ext,omitempty"`               // Placeholder for exchange-specific extensions to OpenRTB.
	Native *Native `json:"native,omitempty"`
}
