package openrtb

import (
	"encoding/json"
	"strconv"
)

type RawString string

func (r RawString) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(r))
}

func (r *RawString) UnmarshalJSON(data []byte) (err error) {
	var s string
	switch data[0] {
	case '"':
		err = json.Unmarshal(data, &s)
	default:
		var i int64
		if err = json.Unmarshal(data, &i); err != nil {
			return err
		}
		s = strconv.FormatInt(i, 10)
	}
	*r = RawString(s)
	return
}

func (r RawString) String() string {
	return string(r)
}
