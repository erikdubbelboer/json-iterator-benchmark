package openrtb

import (
	"encoding/json"
)

type Bid struct {
	ID RawString `json:"id"` // Bidder generated bid ID to assist with logging/tracking.
	// ID    string  `json:"impid"`
	ImpID string  `json:"impid"` // ID of the Imp object in the related bid request.
	Price float64 `json:"price"` // Bid price expressed as CPM although the actual transaction is for a unit impression only. Note that while the type indicates float, integer math is highly recommended when handling currencies (e.g., BigDecimal in Java).

	AdID RawString `json:"adid,omitempty"` // ID of a preloaded ad to be served if the bid wins.
	//	AdID    string   `json:"adid,omitempty"`
	NURL    string   `json:"nurl,omitempty"`    // Win notice URL called by the exchange if the bid wins; optional means of serving ad markup.
	Adm     string   `json:"adm,omitempty"`     // Optional means of conveying ad markup in case the bid wins; supersedes the win notice if markup is included in both.
	Adomain []string `json:"adomain,omitempty"` // Advertiser domain for block list checking (e.g., “ford.com”). This can be an array of for the case of rotating creatives. Exchanges can mandate that only one domain is allowed.
	Bundle  string   `json:"bundle,omitempty"`  // Bundle or package name (e.g., com.foo.mygame) of the app being advertised, if applicable; intended to be a unique ID across exchanges.
	IURL    string   `json:"iurl,omitempty"`    // URL without cache-busting to an image that is representative of the content of the campaign for ad quality/safety checking.

	CID RawString `json:"cid,omitempty"` // Campaign ID to assist with ad quality checking; the collection of creatives for which iurl should be representative.
	//	CID string `json:"cid,omitempty"`

	CrID RawString `json:"crid,omitempty"` // Creative ID to assist with ad quality checking.
	//	CrID string   `json:"crid,omitempty"` // Creative ID to assist with ad quality checking.
	Cat  json.RawMessage `json:"cat,omitempty"`  // IAB content categories of the creative.
	Attr []int           `json:"attr,omitempty"` // Set of attributes describing the creative.

	DealID RawString `json:"deal_id,omitempty"` // Reference to the deal.id from the bid request if this bid pertains to a private marketplace direct deal.
	//	DealID string `json:"deal_id,omitempty"` // Reference to the deal.id from the bid request if this bid pertains to a private marketplace direct deal.
	H   int             `json:"h,omitempty"`   // Height of the creative in pixels.
	W   int             `json:"w,omitempty"`   // Width of the creative in pixels.
	Ext json.RawMessage `json:"ext,omitempty"` // Placeholder for exchange-specific extensions to OpenRTB.
}
