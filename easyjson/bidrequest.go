package openrtb

//go:generate easyjson -pkg -all

type BidRequestExtUdi struct {
	IDFA       string `json:"idfa,omitempty"`
	GAID       string `json:"gaid,omitempty"`
	GoogleAdID string `json:"googleadid,omitempty"` // Smaato
}

type BidRequestExt struct {
	Udi        BidRequestExtUdi `json:"udi,omitempty"`
	RemoteAddr string           `json:"remote_addr,omitempty"`
}

type BidRequest struct {
	ID      RawString     `json:"id"`                // Unique ID of the bid request, provided by the exchange.
	Imp     []Imp         `json:"imp"`               // Array of Imp objects representing the impressions offered. At least 1 Imp object is required.
	Site    *Site         `json:"site,omitempty"`    // Details via a Site object about the publisher’s website. Only applicable and recommended for websites.
	App     *App          `json:"app,omitempty"`     //Details via an App object about the publisher’s app (i.e., non-browser applications). Only applicable and recommended for apps.
	Device  *Device       `json:"device,omitempty"`  // Details via a Device object about the user’s device to which the impression will be delivered.
	User    *User         `json:"user,omitempty"`    // Details via a User object about the human user of the device; the advertising audience.
	Test    int           `json:"test,omitempty"`    // Indicator of test mode in which auctions are not billable, where 0 = live mode, 1 = test mode.
	At      int           `json:"at,omitempty"`      // Auction type, where 1 = First Price, 2 = Second Price Plus. Exchange-specific auction types can be defined using values greater than 500.
	Tmax    int           `json:"tmax,omitempty"`    // Maximum time in milliseconds to submit a bid to avoid timeout. This value is commonly communicated offline.
	Wseat   []string      `json:"wseat,omitempty"`   // Whitelist of buyer seats allowed to bid on this impression. Seat IDs must be communicated between bidders and the exchange a priori. Omission implies no seat restrictions.
	AllImps int           `json:"allimps,omitempty"` // Flag to indicate if Exchange can verify that the impressions offered represent all of the impressions available in context (e.g., all on the web page, all video spots such as pre/mid/post roll) to support road-blocking. 0 = no or unknown, 1 = yes, the impressions offered represent all that are available.
	Cur     []string      `json:"cur,omitempty"`     // Array of allowed currencies for bids on this bid request using ISO-4217 alpha codes. Recommended only if the exchange accepts multiple currencies.
	Bcat    []string      `json:"bcat,omitempty"`    // Blocked advertiser categories using the IAB content categories.
	Badv    []string      `json:"badv,omitempty"`    // Block list of advertisers by their domains (e.g., “ford.com”).
	Regs    *Regs         `json:"regs,omitempty"`    // A Regs object that specifies any industry, legal, or governmental regulations in force for this request.
	Ext     BidRequestExt `json:"ext,omitempty"`
}
