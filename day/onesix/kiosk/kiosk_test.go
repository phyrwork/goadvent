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
		{"aaaaa-bbb-z-y-x",   123, Hash{'a','b','x','y','z'}},
		{"a-b-c-d-e-f-g-h",   987, Hash{'a','b','c','d','e'}},
		{"not-a-real-room",   404, Hash{'o','a','r','e','l'}},
		{"totally-real-room", 200, Hash{'d','e','c','o','y'}},
	}
	got, err := Read(strings.NewReader(in))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("unexpected rooms: want %v, got %v", want, got)
	}
}

func TestDecrypt(t *testing.T) {
	in := Room {
		Name: "qzmt-zixmtkozy-ivhz",
		Sector: 343,
		// Hash not used
	}
	got := Decrypt(in).Name
	want := "very encrypted name"
	if got != want {
		t.Fatalf("unexpected name: want %v, got %v", want, got)
	}
}