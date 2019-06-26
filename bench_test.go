package openrtb

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
	"unsafe"

	"github.com/json-iterator/go"
	"github.com/mailru/easyjson/jwriter"

	easyjson "github.com/erikdubbelboer/json-iterator-benchmark/easyjson"
	normal "github.com/erikdubbelboer/json-iterator-benchmark/normal"
)

var (
	body            = []byte(`{"id":"14854123293616422328","tmax":198,"imp":[{"id":"1","tagid":"361","secure":0,"bidfloor":0.05,"video":{"mimes":["video\/x-flv","video\/mp4","application\/x-shockwave-flash"],"minduration":0,"maxduration":360,"startdelay":0,"linearity":1,"protocols":[1,2,3,7,4,5,6,8],"w":"0","h":"0","delivery":[2],"api":[1,2],"pos":3}}],"site":{"id":"67","page":"http:\/\/www.myrecipes.com","domain":"myrecipes.com","name":"myrecipes.com","mobile":1,"privacypolicy":0,"publisher":{"id":"67"}},"device":{"ua":"Mozilla\/5.0 (Windows NT 6.3; Trident\/7.0; EIE10;ENUSMSN; rv:11.0) like Gecko","ip":"66.222.62.147","language":"en","devicetype":2,"h":-1,"w":-1,"os":"Windows","osv":"8.0","flashver":"-1","ifa":"c74944f941fd5123594ad5e8afcc0557489303c81764e953014149edd707e","didsha1": "60f31729118b6f2f02ab2143e3d057fae77a0fdb","didmd5": "4406046a4560a8ed05a054b546154d8c","dpidsha1": "60f31729118b6f2f02ab2143e3d057fae77a0fdb","dpidmd5": "4406046a4560a8ed05a054b546154d8c","macsha1": "60f31729118b6f2f02ab2143e3d057fae77a0fdb","macmd5": "4406046a4560a8ed05a054b546154d8c","geo":{"country":"USA", "ext":{}},"dnt":1},"user":{"id":"c74944f941fd5123594ad5e8afcc0557489303c81764e953014149edd707e"},"ext":{"ssl":"0","pageviewid":"asmpvx8182801485412329","pagevieworder":"0"}}`)
	normalRequest   = normal.BidRequest{}
	easyjsonRequest = easyjson.BidRequest{}
)

func init() {
	if err := json.Unmarshal(body, &normalRequest); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &easyjsonRequest); err != nil {
		panic(err)
	}
}

func TestUnmarshal(t *testing.T) {
	standard := normal.BidRequest{}
	if err := json.Unmarshal(body, &standard); err != nil {
		t.Fatal(err)
	}

	easy := easyjson.BidRequest{}
	if err := easy.UnmarshalJSON(body); err != nil {
		t.Fatal(err)
	}
	// If we don't do this cast reflect.DeepEqual will immediately say
	// the types are different.
	normalFromEasy := *((*normal.BidRequest)(unsafe.Pointer(&easy)))
	if !reflect.DeepEqual(standard, normalFromEasy) {
		t.Fatal("not the same")
	}

	iter := normal.BidRequest{}
	if err := jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(body, &iter); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(standard, iter) {
		// This test fails because json-iterator doesn't remove whitespace around json.RawMessage
	}
}

func testMarshal(t *testing.T) {
	standard, err := json.Marshal(normalRequest)
	if err != nil {
		t.Fatal(err)
	}

	jw := &jwriter.Writer{}
	easyjsonRequest.MarshalEasyJSON(jw)
	easy, err := jw.BuildBytes()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(standard, easy) {
		t.Fatal("not the same")
	}

	iter, err := jsoniter.Marshal(normalRequest)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(standard, iter) {
		t.Fatal("not the same")
	}
}

func BenchmarkUnmarshalStandard(b *testing.B) {
	request := normal.BidRequest{}

	for i := 0; i < b.N; i++ {
		json.Unmarshal(body, &request)
	}
}

func BenchmarkUnmarshalEasyJson(b *testing.B) {
	request := easyjson.BidRequest{}

	for i := 0; i < b.N; i++ {
		request.UnmarshalJSON(body)
	}
}

func BenchmarkUnmarshalJsonIterator(b *testing.B) {
	request := normal.BidRequest{}

	for i := 0; i < b.N; i++ {
		jsoniter.Unmarshal(body, &request)
	}
}

func BenchmarkUnmarshalJsonIteratorWithEasy(b *testing.B) {
	request := easyjson.BidRequest{}

	for i := 0; i < b.N; i++ {
		jsoniter.Unmarshal(body, &request)
	}
}

func BenchmarkMarshalStandard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(normalRequest)
	}
}

func BenchmarkMarshalEasyJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jw := &jwriter.Writer{}
		easyjsonRequest.MarshalEasyJSON(jw)
		jw.BuildBytes()
	}
}

func BenchmarkMarshalJsonIterator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsoniter.Marshal(normalRequest)
	}
}

func BenchmarkMarshalJsonIteratorWithEasy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsoniter.Marshal(easyjsonRequest)
	}
}
