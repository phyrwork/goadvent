package sqtri

import (
	"reflect"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	tests := map[string]struct {
		in string
		want []Triangle
	}{
		"example": {
			"5 10 25",
			[]Triangle{
				{5, 10, 25},
			},
		},
		"leading spaces": {
			"   5 10 25",
			[]Triangle{
				{5, 10, 25},
			},
		},
		"following newline": {
			"5 10 25\n",
			[]Triangle{
				{5, 10, 25},
			},
		},
		"multiple": {
			"5 10 25\n2 4 9",
			[]Triangle{
				{5, 10, 25},
				{2, 4, 9},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := strings.NewReader(test.in)
			got, err := Read(r)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("unexpected output: want %v, got %v", test.want, got)
			}
		})
	}
}

func TestTriangle_Valid(t *testing.T) {
	tests := map[string]struct {
		tri Triangle
		want bool
	}{
		"lt": {Triangle{10, 5, 25}, false},
		"eq": {Triangle{5, 10, 5}, false},
		"gt": {Triangle{3, 3, 3}, true},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := test.tri.Valid()
			if got != test.want {
				t.Fatalf("unexpected result: want %v, got %v", test.want, got)
			}
		})
	}
}


