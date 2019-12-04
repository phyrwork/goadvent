package kiosk

import (
	"strings"
	"testing"
)

func TestSolveSumReal(t *testing.T) {
	in := strings.Join([]string{
		"aaaaa-bbb-z-y-x-123[abxyz]",
		"a-b-c-d-e-f-g-h-987[abcde]",
		"not-a-real-room-404[oarel]",
		"totally-real-room-200[decoy]",
	}, "\n")
	want := "1514"
	ans := SolveSumReal(strings.NewReader(in))
	if ans.IsError() {
		t.Fatalf("unexpected error: %v", ans)
	}
	if got := ans.String(); got != want {
		t.Fatalf("unexpected answer: want %v, got %v", want, got)
	}
}

