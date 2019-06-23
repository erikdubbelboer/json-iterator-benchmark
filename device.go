package openrtb

import (
	"encoding/json"
)

type DeviceExt struct {
	IDFA string `json:"idfa,omitempty"`
	GAID string `json:"gaid,omitempty"`
}

type Device struct {
	Ua             string          `json:"ua,omitempty"`             // Browser user agent string.
	Geo            *Geo            `json:"geo,omitempty"`            // Location of the device assumed to be the user’s current recommended location defined by a Geo object.
	Dnt            int             `json:"dnt,omitempty"`            // Standard “Do Not Track” flag as set in the header by the recommended browser, where 0 = tracking is unrestricted, 1 = do not track.
	Lmt            int             `json:"lmt,omitempty"`            // “Limit Ad Tracking” signal commercially endorsed (e.g., iOS, recommended Android), where 0 = tracking is unrestricted, 1 = tracking must be limited per commercial guidelines.
	IP             string          `json:"ip,omitempty"`             // IPv4 address closest to device.
	IPv6           string          `json:"ipv6,omitempty"`           // ￼IP address closest to device as IPv6.
	DeviceType     int             `json:"devicetype,omitempty"`     // ￼The general type of device.
	Make           string          `json:"make,omitempty"`           // ￼Device make (e.g., “Apple”).
	Model          string          `json:"model,omitempty"`          //￼ Device model (e.g., “iPhone”).
	OS             string          `json:"os,omitempty"`             // ￼Device operating system (e.g., “iOS”).
	OSv            string          `json:"osv,omitempty"`            //￼ Device operating system version (e.g., “3.1.2”).
	HWv            string          `json:"hwv,omitempty"`            //￼ ￼Hardware version of the device (e.g., “5S” for iPhone 5S).
	H              int             `json:"h,omitempty"`              //￼ Physical height of the screen in pixels.
	W              int             `json:"w,omitempty"`              //￼ Physical width of the screen in pixels.
	PPI            int             `json:"ppi,omitempty"`            //￼ Screen size as pixels per linear inch.
	PxRatio        float64         `json:"pxratio,omitempty"`        //￼ The ratio of physical pixels to device independent pixels.
	JS             int             `json:"js,omitempty"`             //￼ Support for JavaScript, where 0 = no, 1 = yes.
	FlashVer       string          `json:"flashver,omitempty"`       // ￼Version of Flash supported by the browser.
	Language       string          `json:"language,omitempty"`       // ￼Browser language using ISO-639-1-alpha-2.
	Carrier        string          `json:"carrier,omitempty"`        // Carrier or ISP (e.g., “VERIZON”). “WIFI” is often used in mobile to indicate high bandwidth (e.g., video friendly vs. cellular).
	ConnectionType int             `json:"connectiontype,omitempty"` // ￼Network connection type.
	IFA            json.RawMessage `json:"ifa,omitempty"`            // ID sanctioned for advertiser use in the clear (i.e., not hashed).
	DIDSHA1        json.RawMessage `json:"didsha1,omitempty"`        // Hardware device ID (e.g., IMEI); hashed via SHA1.
	DIDMD5         json.RawMessage `json:"didmd5,omitempty"`         // Hardware device ID (e.g., IMEI); hashed via MD5.
	DPIDSHA1       json.RawMessage `json:"dpidsha1,omitempty"`       // Platform device ID (e.g., Android ID); hashed via SHA1.
	DPIDMD5        json.RawMessage `json:"dpidmd5,omitempty"`        // Platform device ID (e.g., Android ID); hashed via MD5.
	MACSHA1        json.RawMessage `json:"macsha1,omitempty"`        // MAC address of the device; hashed via SHA1.
	MACMD5         json.RawMessage `json:"macmd5,omitempty"`         // MAC address of the device; hashed via MD5.
	Ext            DeviceExt       `json:"ext,omitempty"`
}
