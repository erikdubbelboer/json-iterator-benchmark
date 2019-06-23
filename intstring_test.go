package openrtb

import (
	"testing"
)

func TestRawMessageToInt(t *testing.T) {
	var i IntString
	if err := i.UnmarshalJSON([]byte(`123`)); err != nil {
		t.Fatal(err)
	} else if i.Int() != 123 {
		t.Fatalf("expected 123 got %d", i.Int())
	}

	if err := i.UnmarshalJSON([]byte(`"123"`)); err != nil {
		t.Fatal(err)
	} else if i.Int() != 123 {
		t.Fatalf("expected 123 got %d", i.Int())
	}
}
