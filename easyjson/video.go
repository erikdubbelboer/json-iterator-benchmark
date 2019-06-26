package openrtb

type Video struct {
	Mimes          []string  `json:"mimes,omitempty"`          // Content MIME types supported. Popular MIME types may include “video/x-ms-wmv” for Windows Media and “video/x-flv” for Flash Video.
	MinDuration    int       `json:"minduration,omitempty"`    // Minimum video ad duration in seconds.
	MaxDuration    int       `json:"maxduration,omitempty"`    // Maximum video ad duration in seconds.
	Protocol       int       `json:"protocol,omitempty"`       // NOTE: Use of protocols instead is highly recommended. Supported video bid response protocol. At least one supported protocol must be specified in either the protocol or protocols attribute.
	Protocols      []int     `json:"protocols,omitempty"`      // Array of supported video bid response protocols. At least one supported protocol must be specified in either the protocol or protocols attribute.
	W              IntString `json:"w,omitempty"`              // Width of the video player in pixels.
	H              IntString `json:"h,omitempty"`              // Height of the video player in pixels.
	StartDelay     int       `json:"startdelay,omitempty"`     // Indicates the start delay in seconds for pre-roll, mid-roll, or post-roll ad placements.
	Linearity      int       `json:"linearity,omitempty"`      // Indicates if the impression must be linear, nonlinear, etc. If none specified, assume all are allowed.
	Sequence       int       `json:"sequence,omitempty"`       // If multiple ad impressions are offered in the same bid request, the sequence number will allow for the coordinated delivery of multiple creatives.
	Battr          []int     `json:"battr,omitempty"`          // Blocked creative attributes.
	MaxExtended    int       `json:"maxextended,omitempty"`    // Maximum extended video ad duration if extension is allowed. If blank or 0, extension is not allowed. If -1, extension is allowed, and there is no time limit imposed. If greater than 0, then the value represents the number of seconds of extended play supported beyond the maxduration value.
	MinBitrate     int       `json:"minbitrate,omitempty"`     // Minimum bit rate in Kbps. Exchange may set this dynamically or universally across their set of publishers.
	MaxBitrate     int       `json:"maxbitrate,omitempty"`     // Maximum bit rate in Kbps. Exchange may set this dynamically or universally across their set of publishers.
	BoxingAllowed  int       `json:"boxingallowed,omitempty"`  // Indicates if letter-boxing of 4:3 content into a 16:9 window is allowed, where 0 = no, 1 = yes.
	PlaybackMethod []int     `json:"playbackmethod,omitempty"` // Allowed playback methods. If none specified, assume all are allowed.
	Delivery       []int     `json:"delivery,omitempty"`       // Supported delivery methods (e.g., streaming, progressive). If none specified, assume all are supported.
	Pos            int       `json:"pos,omitempty"`            // Ad position on screen.
	CompanionAd    []Banner  `json:"companionad,omitempty"`    // Array of Banner objects if companion ads are available.
	API            []int     `json:"api,omitempty"`            // List of supported API frameworks for this impression. If an API is not explicitly listed, it is assumed not to be supported.
	CompanionType  []int     `json:"companiontype,omitempty"`  // Supported VAST companion ad types. Recommended if companion Banner objects are included via the companionad array.
	//Ext            Extension `json:"ext,omitempty"`            // Placeholder for exchange-specific extensions to OpenRTB.
}
