package openrtb

type BidResponse struct {
	ID         RawString `json:"id"`                   // ID of the bid request to which this is a response.
	SeatBid    []SeatBid `json:"seatbid,omitempty"`    // Array of seatbid objects; 1+ required if a bid is to be made.
	BidID      string    `json:"bidid,omitempty"`      // Bidder generated response ID to assist with logging/tracking.
	Cur        string    `json:"cur,omitempty"`        // Bid currency using ISO-4217 alpha codes.
	CustomData string    `json:"customdata,omitempty"` // Optional feature to allow a bidder to set data in the exchange’s cookie. The string must be in base85 cookie safe characters and be in any format. Proper JSON encoding must be used to include “escaped” quotation marks.
	NBR        int       `json:"nbr,omitempty"`        // Reason for not bidding.
	//Ext     Extension `json:"ext,omitempty"`     // Placeholder for exchange-specific extensions to OpenRTB.
}
