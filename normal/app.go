package openrtb

import (
	"encoding/json"
)

type App struct {
	ID            RawString       `json:"id,omitempty"`            // Exchange-specific app ID.
	Name          string          `json:"name,omitempty"`          // App name (may be aliased at the publisher’s request).
	Bundle        string          `json:"bundle,omitempty"`        // Application bundle or package name (e.g., com.foo.mygame); intended to be a unique ID across exchanges.
	Domain        string          `json:"domain,omitempty"`        // Domain of the app (e.g., “mygame.foo.com”).
	StoreURL      json.RawMessage `json:"storeurl,omitempty"`      // App store URL for an installed app; for QAG 1.5 compliance.
	Cat           json.RawMessage `json:"cat,omitempty"`           // Array of IAB content categories of the app.
	SectionCat    []string        `json:"sectioncat,omitempty"`    // Array of IAB content categories that describe the current section of the app.
	PageCat       []string        `json:"pagecat,omitempty"`       // Array of IAB content categories that describe the current page or view of the app.
	Ver           json.RawMessage `json:"ver,omitempty"`           // Application version.
	PrivacyPolicy int             `json:"privacypolicy,omitempty"` // Indicates if the app has a privacy policy, where 0 = no, 1 = yes.
	Paid          int             `json:"paid,omitempty"`          // 0 = app is free, 1 = the app is a paid version.
	Publisher     *Publisher      `json:"publisher,omitempty"`     // Details about the Publisher of the app.
	Content       *Content        `json:"content,omitempty"`       // Details about the Content within the app.
	Keywords      string          `json:"keywords,omitempty"`      // Comma separated list of keywords about the app.
	//Ext            Extension `json:"ext,omitempty"`             // Placeholder for exchange-specific extensions to OpenRTB.
}
