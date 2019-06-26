package openrtb

type Geo struct {
	Lat           float64 `json:"lat,omitempty"`           // Latitude from -90.0 to +90.0, where negative is south.
	Lon           float64 `json:"lon,omitempty"`           // Longitude from -180.0 to +180.0, where negative is west.
	Type          int     `json:"type,omitempty"`          // Source of location data; recommended when passing lat/lon.
	Country       string  `json:"country,omitempty"`       // Country code using ISO-3166-1-alpha-3.
	Region        string  `json:"region,omitempty"`        // Region code using ISO-3166-2; 2-letter state code if USA.
	RegionFIPS104 string  `json:"regionfips104,omitempty"` // Region of a country using FIPS 10-4 notation. While OpenRTB supports this attribute, it has been withdrawn by NIST in 2008.
	Metro         string  `json:"metro,omitempty"`         // Google metro code; similar to but not exactly Nielsen DMAs.
	City          string  `json:"city,omitempty"`          // City using United Nations Code for Trade & Transport Locations.
	Zip           string  `json:"zip,omitempty"`           // Zip or postal code.
	UTCOffset     int     `json:"utcoffset,omitempty"`     // Local device time as the number +/- of minutes from UTC.
	//Ext           Ext     `json:"ext,omitempty"`           // Placeholder for exchange-specific extensions to OpenRTB.
}
