package realloc

import (
	"reflect"
	"strings"
	"testing"
)

func TestScanMemory(t *testing.T) {
	tests := []struct {
		s string
		m Memory
	}{
		{"0 2 7 0", NewMemory(0, 2, 7, 0)},
		{
			"4	1	15	12	0	9	9	5	5	8	7	3	14	5	12	3",
			NewMemory(4, 1, 15, 12, 0, 9, 9, 5, 5, 8, 7, 3, 14, 5, 12, 3),
		},
	}
	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			r := strings.NewReader(test.s)
			m, err := ScanMemory(r)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(m, test.m) {
				t.Fatalf("unexpected value: want %#v, got %#v", test.m, m)
			}
		})
	}
}

func TestRealloc(t *testing.T) {
	tests := []struct {
		b, a Memory // before and after
	}{
		{NewMemory(0, 2, 7, 0), NewMemory(2, 4, 1, 2)},
		{NewMemory(2, 4, 1, 2), NewMemory(3, 1, 2, 3)},
		{NewMemory(3, 1, 2, 3), NewMemory(0, 2, 3, 4)},
		{NewMemory(0, 2, 3, 4), NewMemory(1, 3, 4, 1)},
		{NewMemory(1, 3, 4, 1), NewMemory(2, 4, 1, 2)},
	}
	for _, test := range tests {
		name := test.b.String()
		t.Run(name, func(t *testing.T) {
			Realloc(test.b)
			if !reflect.DeepEqual(test.b, test.a) {
				t.Fatalf("unexpected value: want %#v, got %#v", test.a, test.b)
			}
		})
	}
}

func TestReallocUniq(t *testing.T) {
	m := NewMemory(0, 2, 7, 0)
	c := ReallocUniq(m)
	e := 5
	if c != e {
		t.Fatalf("unexpected value: want %v, got %v", e, c)
	}
}

