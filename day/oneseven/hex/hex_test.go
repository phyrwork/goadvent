package hex

import (
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := map[string]int {
		"ne,ne,ne": 3,
		"ne,ne,sw,sw": 0,
		"ne,ne,s,s": 2,
		"se,sw,se,sw,sw": 3,
	}
	for s, want := range tests {
		t.Run(s, func(t *testing.T) {
			r := strings.NewReader(s)
			got, err := Solve(r)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != want {
				t.Fatalf("unexpected value: want %v, got %v", want, got)
			}
		})
	}
}
