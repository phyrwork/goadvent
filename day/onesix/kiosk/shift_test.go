package kiosk

import (
	"reflect"
	"testing"
)

// qzmt-zixmtkozy-ivhz-343

func TestShifters(t *testing.T) {
	// an all-in-one test to validate my weird mix of laziness and
	// over-engineering
	in := []rune("qzmt-zixmtkozy-ivhz")
	ws := WrapShifter{343, 'a', 'z'}
	ms := MaskShifter{ws, func (_ int, c rune) bool { return c == '-' }}
	want := []rune("very-encrypted-name")
	got := ms.Shift(in)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("unexpected output: want %v, got %v", string(want), string(got))
	}
}

