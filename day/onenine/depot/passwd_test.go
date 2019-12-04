package depot

import (
	"strconv"
	"testing"
)

func TestLengthRule(t *testing.T) {
	tests := map[int]struct {
		s string
		n int
		ok bool
	}{
		1: {"12345", 6, false},
		2: {"qwerty", 6, true},
		3: {"lololol", 6, false},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ok := NewLengthRule(test.n)(test.s)
			if ok != test.ok {
				t.Fatalf("unexpected result: want: %v, got %v", test.ok, ok)
			}
		})
	}
}

func TestAdjacentLeastRule(t *testing.T) {
	tests := map[int]struct {
		s string
		n int
		ok bool
	}{
		1: {"12345", 2, false},
		2: {"qwerty", 2, false},
		3: {"lololol", 2, false},
		4: {"halloome", 2, true},
		5: {"oops", 2, true},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ok := NewAdjacentLeastRule(test.n)(test.s)
			if ok != test.ok {
				t.Fatalf("unexpected result: want: %v, got %v", test.ok, ok)
			}
		})
	}
}

func TestAdjacentExactRule(t *testing.T) {
	tests := map[int]struct {
		s string
		n int
		ok bool
	}{
		1: {"12345", 2, false},
		2: {"qwerty", 2, false},
		3: {"lololol", 2, false},
		4: {"halloome", 2, true},
		5: {"oops", 2, true},
		6: {"112233", 2, true},
		7: {"123444", 2, false},
		8: {"111122", 2, true},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ok := NewAdjacentExactRule(test.n)(test.s)
			if ok != test.ok {
				t.Fatalf("unexpected result: want: %v, got %v", test.ok, ok)
			}
		})
	}
}

func TestDecreaseRule(t *testing.T) {
	tests := map[int]struct {
		s string
		ok bool
	}{
		1: {"12345", true},
		2: {"54321", false},
		3: {"243779", false},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ok := DecreaseRule(test.s)
			if ok != test.ok {
				t.Fatalf("unexpected result: want: %v, got %v", test.ok, ok)
			}
		})
	}
}
