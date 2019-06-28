package passwd

import "testing"

func TestUniqWords(t *testing.T) {
	tests := map[Passphrase]bool {
		"aa bb cc dd ee": true,
		"aa bb cc dd aa": false,
		"aa bb cc dd aaa": true,
	}
	for p, want := range tests {
		t.Run(string(p), func(t *testing.T) {
			got, err := UniqWords(p)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != want {
				t.Fatalf("unexpected result: want %v, got %v", want, got)
			}
		})
	}
}

func TestUniqAnagrams(t *testing.T) {
	tests := map[Passphrase]bool {
		"abcde fghij": true,
		"abcde xyz ecdab": false,
		"a ab abc abd abf abj": true,
		"iiii oiii ooii oooi oooo": true,
		"oiii ioii iioi iiio": false,
	}
	for p, want := range tests {
		t.Run(string(p), func(t *testing.T) {
			got, err := UniqAnagrams(p)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != want {
				t.Fatalf("unexpected result: want %v, got %v", want, got)
			}
		})
	}
}

