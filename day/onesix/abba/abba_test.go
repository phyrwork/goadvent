package abba

import (
	"reflect"
	"testing"
)

func TestSequencer(t *testing.T) {
	tests := map[Address][]Sequence {
		"abba[mnop]qrst": {
			{'a', 'b', 'b', 'a'},
			{'m', 'n', 'o', 'p'},
			{'q', 'r', 's', 't'},
		},
		"ioxxoj[asdfgh]zxcvbn": {
			{'i', 'o', 'x', 'x'},
			{'o', 'x', 'x', 'o'},
			{'x', 'x', 'o', 'j'},
			{'a', 's', 'd', 'f'},
			{'s', 'd', 'f', 'g'},
			{'d', 'f', 'g', 'h'},
			{'z', 'x', 'c', 'v'},
			{'x', 'c', 'v', 'b'},
			{'c', 'v', 'b', 'n'},
		},
	}
	for addr, want := range tests {
		t.Run(string(addr), func(t *testing.T) {
			got := make([]Sequence, 0)
			seq := NewSequencer(addr)
			for seq.Next() {
				got = append(got, seq.Sequence())
			}
			if !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected sequences: want %v, got %v", want, got)
			}
		})
	}
}


