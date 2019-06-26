package openrtb

import (
	"encoding/json"
)

type Publisher struct {
	ID     RawString       `json:"id,omitempty"`     // Exchange-specific publisher ID.
	Name   string          `json:"name,omitempty"`   // Publisher name (may be aliased at the publisher’s request).
	Cat    json.RawMessage `json:"cat,omitempty"`    // Array of IAB content categories that describe the publisher.
	Domain string          `json:"domain,omitempty"` // Highest level domain of the publisher (e.g., “publisher.com”).
	//Ext            Extension `json:"ext,omitempty"`             // Placeholder for exchange-specific extensions to OpenRTB.
}
