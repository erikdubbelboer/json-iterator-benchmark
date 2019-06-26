package openrtb

import (
	"encoding/json"
)

type User struct {
	ID         string          `json:"id,omitempty"`         // Exchange-specific ID for the user. At least one of id or buyerid is recommended.
	BuyerUID   string          `json:"buyeruid,omitempty"`   // Buyer-specific ID for the user as mapped by the exchange for the buyer. At least one of buyerid or id is recommended.
	YOB        json.RawMessage `json:"yob,omitempty"`        // Year of birth as a 4-digit integer.
	Gender     string          `json:"gender,omitempty"`     // Gender, where “M” = male, “F” = female, “O” = known to be other (i.e., omitted is unknown).
	Keywords   string          `json:"keywords,omitempty"`   // Comma separated list of keywords, interests, or intent.
	CustomData string          `json:"customdata,omitempty"` // Optional feature to pass bidder data that was set in the exchange’s cookie. The string must be in base85 cookie safe characters and be in any format. Proper JSON encoding must be used to include “escaped” quotation marks.
	Geo        *Geo            `json:"geo,omitempty"`        // Location of the user’s home base defined by a Geo object. This is not necessarily their current location.
	Data       []Data          `json:"data,omitempty"`       // Additional user data. Each Data object represents a different data source.
	Ext        Ext             `json:"ext,omitempty"`        // Placeholder for exchange-specific extensions to OpenRTB.
}

type Data struct {
	ID      string    `json:"id,omitempty"`      // Exchange-specific ID for the data provider.
	Name    string    `json:"name,omitempty"`    // Exchange-specific name for the data provider.
	Segment []Segment `json:"segment,omitempty"` // Array of Segment objects that contain the actual data values.
	//Ext        Ext    `json:"ext,omitempty"`        // Placeholder for exchange-specific extensions to OpenRTB.
}

type Segment struct {
	ID    string `json:"id,omitempty"`    // ID of the data segment specific to the data provider.
	Name  string `json:"name,omitempty"`  // Name of the data segment specific to the data provider.
	Value string `json:"value,omitempty"` // String representation of the data segment value.
	//Ext        Ext    `json:"ext,omitempty"`        // Placeholder for exchange-specific extensions to OpenRTB.
}
