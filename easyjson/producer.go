package openrtb

import (
	"encoding/json"
)

type Producer struct {
	ID     string          `json:"id,omitempty"`     // Content producer or originator ID. Useful if content is syndicated and may be posted on a site using embed tags.
	Name   string          `json:"name,omitempty"`   // Content producer or originator name (e.g., “Warner Bros”).
	Cat    json.RawMessage `json:"cat,omitempty"`    // Array of IAB content categories that describe the content producer.
	Domain string          `json:"domain,omitempty"` // Highest level domain of the content producer (e.g., “producer.com”).
	//Ext    Extension `json:"ext,omitempty"`    // Placeholder for exchange-specific extensions to OpenRTB.
}
