package taxi

import (
	"github.com/phyrwork/goadvent/vector"
	"reflect"
	"strings"
	"testing"
)

var tests = map[string]struct {
	in    string
	steps []Step
	comp  Compass
	end   vector.Vector
	rpt   vector.Vector
}{
	"example 1": {
		"R2, L3\n",
		[]Step{
			{Right, 2},
			{Left, 3},
		},
		Compass(North),
		vector.Vector{2, 3},
		nil,
	},
	"example 2": {
		"R2, R2, R2\n",
		[]Step{
			{Right, 2},
			{Right, 2},
			{Right, 2},
		},
		Compass(North),
		vector.Vector{0, -2},
		nil,
	},
	"example 3": {
		"R5, L5, R5, R3\n",
		[]Step{
			{Right, 5},
			{Left, 5},
			{Right, 5},
			{Right, 3},
		},
		Compass(North),
		vector.Vector{10, 2},
		nil,
	},
	"example 4": {
		"R8, R4, R4, R8\n",
		[]Step{
			{Right, 8},
			{Right, 4},
			{Right, 4},
			{Right, 8},
		},
		Compass(North),
		vector.Vector{4, 4},
		vector.Vector{4, 0},
	},
}

func TestRead(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := strings.NewReader(test.in)
			got, err := Read(r)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(test.steps, got) {
				t.Fatalf("unexpected value: want %v, got %v", test.steps, got)
			}
		})
	}
}

func TestWalk(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			c := Compass(North)
			got := Walk(c, test.steps...)
			if !reflect.DeepEqual(got, test.end) {
				t.Fatalf("unexpected vector: want %v, got %v", test.end, got)
			}
		})
	}
}

func TestWalkUntilRevisit(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			c := Compass(North)
			got := WalkUntilRevisit(c, test.steps...)
			if !reflect.DeepEqual(got, test.rpt) {
				t.Fatalf("unexpected vector: want %v, got %v", test.rpt, got)
			}
		})
	}
}

