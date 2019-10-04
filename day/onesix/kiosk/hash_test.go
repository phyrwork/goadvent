package kiosk

import "testing"

func TestAlphaCountHash(t *testing.T) {
	tests := map[string]Hash {
		"aaaaa-bbb-z-y-x": {'a','b','x','y','z'},
		"a-b-c-d-e-f-g-h": {'a','b','c','d','e'},
		"not-a-real-room-404": {'o','a','r','e','l'},
		"totally-real-room-200": {'l','o','a','r','t'},
	}
	for in, want := range tests {
		t.Run(in, func(t *testing.T) {
			got, err := AlphaCountHash(in)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != want {
				t.Fatalf("unexpected hash: want %v, got %v", string(want[:]), string(got[:]))
			}
		})
	}
}

