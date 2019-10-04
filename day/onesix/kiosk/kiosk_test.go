package kiosk

import (
	"reflect"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	in := strings.Join([]string{
		"aaaaa-bbb-z-y-x-123[abxyz]",
		"a-b-c-d-e-f-g-h-987[abcde]",
		"not-a-real-room-404[oarel]",
		"totally-real-room-200[decoy]",
	}, "\n")
	want := []Room{
		{"aaaaa-bbb-z-y-x",   123, "abxyz"},
		{"a-b-c-d-e-f-g-h",   987, "abcde"},
		{"not-a-real-room",   404, "oarel"},
		{"totally-real-room", 200, "decoy"},
	}
	got, err := Read(strings.NewReader(in))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("unexpected rooms: want %v, got %v", want, got)
	}
}

