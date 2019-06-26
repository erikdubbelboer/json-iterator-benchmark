package openrtb

import (
	"encoding/json"
)

type Content struct {
	ID                 RawString       `json:"id,omitempty"`                 // ￼ID uniquely identifying the content.
	Episode            int             `json:"episode,omitempty"`            // ￼Episode number (typically applies to video content).
	Title              string          `json:"title,omitempty"`              // Content title. Video Examples: “Search Committee” (television), “A New Hope” (movie), or “Endgame” (made for web). Non-Video Example: “Why an Antarctic Glacier Is Melting So Quickly” (Time magazine article).
	Series             string          `json:"series,omitempty"`             // Content series. Video Examples: “The Office” (television), “Star Wars” (movie), or “Arby ‘N’ The Chief” (made for web). Non-Video Example: “Ecocentric” (Time Magazine blog).
	Season             string          `json:"season,omitempty"`             // ￼Content season; typically for video content (e.g., “Season 3”).
	Producer           *Producer       `json:"producer,omitempty"`           // Details about the content Producer
	URL                string          `json:"url,omitempty"`                // ￼URL of the content, for buy-side contextualization or review.
	Cat                json.RawMessage `json:"cat,omitempty"`                // Array of IAB content categories that describe the content producer.
	VideoQuality       int             `json:"videoquality,omitempty"`       //￼ Video quality per IAB’s classification.
	Context            int             `json:"context,omitempty"`            //￼ Type of content (game, video, text, etc.).
	ContentRating      string          `json:"contentrating,omitempty"`      //￼ Content rating (e.g., MPAA).
	UserRating         string          `json:"userrating,omitempty"`         //￼ User rating of the content (e.g., number of stars, likes, etc.).
	QAGMediaRating     int             `json:"qagmediarating,omitempty"`     //￼ Media rating per QAG guidelines.
	Keywords           string          `json:"keywords,omitempty"`           //￼ Comma separated list of keywords describing the content.
	LiveStream         int             `json:"livestream,omitempty"`         // ￼0 = not live, 1 = content is live (e.g., stream, live blog).
	SourceRelationship int             `json:"sourcerelationship,omitempty"` //￼ ￼0 = indirect, 1 = direct.
	Len                int             `json:"len,omitempty"`                //￼ Length of content in seconds; appropriate for video or audio.
	Language           string          `json:"language,omitempty"`           // Content language using ISO-639-1-alpha-2.
	Embeddable         int             `json:"embeddable,omitempty"`         // Indicator of whether or not the content is embeddable (e.g., an embeddable video player), where 0 = no, 1 = yes.
	//Ext                Extension `json:"ext,omitempty"`                // Placeholder for exchange-specific extensions to OpenRTB.
}
