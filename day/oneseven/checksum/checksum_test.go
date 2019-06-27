package checksum

import (
	"github.com/phyrwork/goadvent/iterator"
	"strings"
	"testing"
)

//

var example1 = `5 1 9 5
7 5 3
2 4 6 8`

var example2 = `5 9 2 8
9 4 7 3
3 8 6 5`

var rows = map[string]struct{
	it      iterator.ResetIterator
	min     int
	max     int
	diff    int
	factdiv int
}{
	"1": {iterator.NewArrayIterator(5, 1, 9, 5), 1, 9, 8, 0},
	"2": {iterator.NewArrayIterator(7, 5, 3),    3, 7, 4, 0},
	"3": {iterator.NewArrayIterator(2, 4, 6, 8), 2, 8, 6, 0},
	"4": {iterator.NewArrayIterator(5, 9, 2, 8), 2, 9, 7, 4},
	"5": {iterator.NewArrayIterator(9, 4, 7, 3), 3, 9, 6, 3},
	"6": {iterator.NewArrayIterator(3, 8, 6, 5), 3, 8, 5, 2},
}

var pages = map[string]struct{
	it  iterator.ResetIterator
	sum SumFunc
	chk int
}{
	"1 (numeric)": {
		iterator.NewArrayIterator(
			iterator.NewArrayIterator(5, 1, 9, 5),
			iterator.NewArrayIterator(7, 5, 3),
			iterator.NewArrayIterator(2, 4, 6, 8),
		),
		Diff,
		18,
	},
	"1 (text)": {
		NewRowScanner(strings.NewReader(example1)),
		Diff,
		18,
	},
	"2 (text)": {
		NewRowScanner(strings.NewReader(example2)),
		FactorDiv,
		9,
	},
}

func TestMin(t *testing.T) {
	for name, test := range rows {
		t.Run(name, func(t *testing.T) {
			if err := test.it.Reset(); err != nil {
				 t.Fatalf("unexpected error: %v", err)
			}
			r, err := Min(test.it)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if r != test.min {
				t.Fatalf("unexpected value: want %v, got %v", test.min, r)
			}
		})
	}
}

func TestMax(t *testing.T) {
	for name, test := range rows {
		t.Run(name, func(t *testing.T) {
			if err := test.it.Reset(); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			r, err := Max(test.it)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if r != test.max {
				t.Fatalf("unexpected value: want %v, got %v", test.max, r)
			}
		})
	}
}

func TestDiff(t *testing.T) {
	for name, test := range rows {
		t.Run(name, func(t *testing.T) {
			if err := test.it.Reset(); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			r, err := Diff(test.it)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if r != test.diff {
				t.Fatalf("unexpected value: want %v, got %v", test.diff, r)
			}
		})
	}
}

func TestFactorDiv(t *testing.T) {
	for name, test := range rows {
		t.Run(name, func(t *testing.T) {
			if err := test.it.Reset(); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			r, err := FactorDiv(test.it)
			if test.factdiv > 0 {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if r != test.factdiv {
					t.Fatalf("unexpected value: want %v, got %v", test.factdiv, r)
				}
			} else {
				if err == nil {
					t.Fatalf("expected error")
				}
			}
		})
	}
}

func TestChecksum(t *testing.T) {
	for name, test := range pages {
		t.Run(name, func(t *testing.T) {
			if err := test.it.Reset(); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			r, err := Checksum(test.it, test.sum)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if r != test.chk {
				t.Fatalf("unexpected value: want %v, got %v", test.chk, r)
			}
		})
	}
}



