package houses

import (
	"github.com/phyrwork/goadvent/iterator"
	"strings"
	"testing"
)

func TestUnique(t *testing.T) {
	tests := map[string]int {
		">": 2,
		"^>v<": 4,
		"^v^v^v^v^v": 2,
	}
	for in, want := range tests {
		t.Run(in, func(t *testing.T) {
			r := strings.NewReader(in)
			it := iterator.NewRuneScanner(r)
			got, err := Solve(it)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != want {
				t.Fatalf("unexpected value: want %v, got %v", want, got)
			}
		})
	}
}

