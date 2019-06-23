package openrtb

import (
	"encoding/json"
	"strconv"
)

type IntString int

func (i IntString) MarshalJSON() (r []byte, err error) {
	strconv.AppendInt(r, int64(i), 10)
	return
}

func (i *IntString) UnmarshalJSON(data []byte) (err error) {
	var r int
	switch data[0] {
	case '"':
		r, err = strconv.Atoi(string(data[1 : len(data)-1]))
	default:
		err = json.Unmarshal(data, &r)
	}
	*i = IntString(r)
	return
}

func (i IntString) Int() int {
	return int(i)
}
