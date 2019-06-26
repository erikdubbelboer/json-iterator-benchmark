package openrtb

type Banner struct {
	W        int       `json:"w"`                  // Width of the impression in pixels. If neither wmin nor wmax are specified, this value is an exact width requirement. Otherwise it is a preferred width.
	H        int       `json:"h"`                  // Height of the impression in pixels. If neither hmin nor hmax are specified, this value is an exact height requirement. Otherwise it is a preferred height.
	Wmax     int       `json:"wmax,omitempty"`     // Maximum width of the impression in pixels. If included along with a w value then w should be interpreted as a recommended or preferred width.
	Hmax     int       `json:"hmax,omitempty"`     // Maximum height of the impression in pixels. If included along with an h value then h should be interpreted as a recommended or preferred height.
	Wmin     int       `json:"wmin,omitempty"`     // Minimum width of the impression in pixels. If included along with a w value then w should be interpreted as a recommended or preferred width.
	Hmin     int       `json:"hmin,omitempty"`     // Minimum height of the impression in pixels. If included along with an h value then h should be interpreted as a recommended or preferred height.
	ID       RawString `json:"id,omitempty"`       // Unique identifier for this banner object. Recommended when Banner objects are used with a Video object to represent an array of companion ads. Values usually start at 1 and increase with each object; should be unique within an impression.
	Btype    []int     `json:"btype,omitempty"`    // Blocked banner ad types.
	Battr    []int     `json:"battr,omitempty"`    // Blocked creative attributes.
	Pos      int       `json:"pos,omitempty"`      // Ad position on screen.
	Mimes    []string  `json:"mimes,omitempty"`    // Content MIME types supported. Popular MIME types may include “application/x-shockwave-flash”, “image/jpg”, and “image/gif”.
	TopFrame int       `json:"topframe,omitempty"` // Indicates if the banner is in the top frame as opposed to an iframe, where 0 = no, 1 = yes.
	ExpDir   []int     `json:"expdir,omitempty"`   // Directions in which the banner may expand.
	API      []int     `json:"api,omitempty"`      // List of supported API frameworks for this impression. If an API is not explicitly listed, it is assumed not to be supported.
	//Ext      Extension `json:"ext,omitempty"`      // Placeholder for exchange-specific extensions to OpenRTB.
}
