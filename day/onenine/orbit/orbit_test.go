package orbit

import (
	"reflect"
	"strings"
	"testing"
)

func TestReadOrbits(t *testing.T) {
	s := strings.Join([]string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
	}, "\n")
	r := strings.NewReader(s)
	got, err := ReadOrbits(r)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := []Orbit{
		{"B", "COM"},
		{"C", "B"},
		{"D", "C"},
		{"E", "D"},
		{"F", "E"},
		{"G", "B"},
		{"H", "G"},
		{"I", "D"},
		{"J", "E"},
		{"K", "J"},
		{"L", "K"},
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("unexpected orbits: want %v, got %v", want, got)
	}
}

func TestCountOrbits(t *testing.T) {
	o := []Orbit{
		{"B", "COM"},
		{"C", "B"},
		{"D", "C"},
		{"E", "D"},
		{"F", "E"},
		{"G", "B"},
		{"H", "G"},
		{"I", "D"},
		{"J", "E"},
		{"K", "J"},
		{"L", "K"},
	}
	g := NewGraph(o...)
	got := CountOrbits(g)
	want := 42
	if got != want {
		t.Fatalf("unexpected count: want %v, got %v", want, got)
	}
}

