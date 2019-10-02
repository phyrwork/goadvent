package keypad

import (
	"reflect"
	"strings"
	"testing"
)

func TestDecoder(t *testing.T) {
	// with default grammar, keypad and cursor
	dc := NewDecoder(DefaultGrammar, DefaultKeypad, DefaultPosition)
	// with example
	inst := "ULL\nRRDDD\nLURDL\nUUUUD"
	// go!
	r := strings.NewReader(inst)
	if err := dc.Decode(r); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []int{1,9,8,5}
	if got := dc.Out(); !reflect.DeepEqual(want, got) {
		t.Fatalf("unexpected out: want %v, got %v", want, got)
	}
}

