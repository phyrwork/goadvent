package sqtri

import (
	"reflect"
	"strings"
	"testing"
)

func TestReadRows(t *testing.T) {
	tests := map[string]struct {
		in string
		want Triangles
	}{
		"example": {
			"5 10 25",
			Triangles{
				{5, 10, 25},
			},
		},
		"leading spaces": {
			"   5 10 25",
			Triangles{
				{5, 10, 25},
			},
		},
		"following newline": {
			"5 10 25\n",
			Triangles{
				{5, 10, 25},
			},
		},
		"multiple": {
			"5 10 25\n2 4 9",
			Triangles{
				{5, 10, 25},
				{2, 4, 9},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := strings.NewReader(test.in)
			got, err := ReadRows(r)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("unexpected output: want %v, got %v", test.want, got)
			}
		})
	}
}

func TestReadTrios(t *testing.T) {
	tests := map[string]struct {
		in string
		want Triangles
	}{
		"example": {
			"101 301 501\n102 302 502\n103 303 503\n201 401 601\n202 402 602\n203 403 603",
			Triangles{
				{101, 102, 103},
				{301, 302, 303},
				{501, 502, 503},
				{201, 202, 203},
				{401, 402, 403},
				{601, 602, 603},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := strings.NewReader(test.in)
			got, err := ReadTrios(r)
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


