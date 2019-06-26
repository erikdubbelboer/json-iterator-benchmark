package openrtb

import (
	"encoding/json"
)

type Site struct {
	ID            RawString       `json:"id,omitempty"`            // Exchange-specific site ID.
	Name          string          `json:"name,omitempty"`          // Site name (may be aliased at the publisher’s request).
	Domain        string          `json:"domain,omitempty"`        // Domain of the site (e.g., “mysite.foo.com”).
	Cat           json.RawMessage `json:"cat,omitempty"`           // Array of IAB content categories of the site.
	SectionCat    []string        `json:"sectioncat,omitempty"`    // Array of IAB content categories that describe the current section of the site.
	PageCat       []string        `json:"pagecat,omitempty"`       // Array of IAB content categories that describe the current page or view of the site.
	Page          string          `json:"page,omitempty"`          // URL of the page where the impression will be shown.
	Ref           string          `json:"ref,omitempty"`           // Referrer URL that caused navigation to the current page.
	Search        string          `json:"search,omitempty"`        // Search string that caused navigation to the current page.
	Mobile        int             `json:"mobile,omitempty"`        // Mobile-optimized signal, where 0 = no, 1 = yes.
	PrivacyPolicy int             `json:"privacypolicy,omitempty"` // Indicates if the site has a privacy policy, where 0 = no, 1 = yes.
	Publisher     *Publisher      `json:"publisher,omitempty"`     // Details about the Publisher of the site.
	Content       *Content        `json:"content,omitempty"`       // Details about the Content within the site.
	Keywords      string          `json:"keywords,omitempty"`      // Comma separated list of keywords about the site.
	//Ext            Extension `json:"ext,omitempty"`             // Placeholder for exchange-specific extensions to OpenRTB.
}
