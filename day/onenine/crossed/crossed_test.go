package crossed

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	tests := map[int]struct {
		in string
		want [][]Vector
	}{
		1: {
			strings.Join([]string{
				"R75,D30,R83,U83,L12,D49,R71,U7,L72",
				"U62,R66,U55,R34,D71,R55,D58,R83",
			}, "\n"),
			[][]Vector{
				{{75, 0}, {0, -30}, {83, 0}, {0, 83}, {-12, 0}, {0, -49}, {71, 0}, {0, 7}, {-72, 0}},
				{{0, 62}, {66, 0}, {0, 55}, {34, 0}, {0, -71}, {55, 0}, {0, -58}, {83, 0}},
			},
		},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := Read(strings.NewReader(test.in))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("unexpected vectors: want %v, got %v", test.want, got)
			}
		})
	}
}

func TestVectorLineBuilder(t *testing.T) {
	tests := map[int]struct {
		in []Vector
		from Point
		want Line
	}{
		// R8,U5,L5,D3
		1: {
			[]Vector{{8, 0}, {0, 5}, {-5, 0}, {0, -3}},
			Point{0, 0},
			Line{
				Segment{Point{0, 0}, Point{8, 0}},
				Segment{Point{8, 0}, Point{8, 5}},
				Segment{Point{8, 5}, Point{3, 5}},
				Segment{Point{3, 5}, Point{3, 2}},
			},
		},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := VectorLineBuilder{test.from}
			got := Segments(b.New(test.in...)...)
			want := Segments(test.want...)
			if !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected segments: want %v, got %v", want, got)
			}
		})
	}
}

func TestLine_Map(t *testing.T) {
	tests := map[int]struct {
		in Line
		want Map
	}{
		1: {
			Line{
				Segment{Point{0, 0}, Point{8, 0}},
				Segment{Point{8, 0}, Point{8, 5}},
				Segment{Point{8, 5}, Point{3, 5}},
				Segment{Point{3, 5}, Point{3, 2}},
			},
			Map{
					{0, 0}: {},
					{1, 0}: {},
					{2, 0}: {},
					{3, 0}: {},
					{4, 0}: {},
					{5, 0}: {},
					{6, 0}: {},
					{7, 0}: {},
					{8, 0}: {},
					{8, 1}: {},
					{8, 2}: {},
					{8, 3}: {},
					{8, 4}: {},
					{8, 5}: {},
					{7, 5}: {},
					{6, 5}: {},
					{5, 5}: {},
					{4, 5}: {},
					{3, 5}: {},
					{3, 4}: {},
					{3, 3}: {},
					{3, 2}: {},
			},
		},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := test.in.Map()
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("unexpected map: want %v, got %v", test.want, got)
			}
		})
	}
}

func TestMap_Intersect(t *testing.T) {
	tests := map[int]struct {
		a, b Map
		want Map
	}{
		1: {
			Map{
				{0, 1}: {},
				{0, 2}: {},
				{0, 3}: {},
			},
			Map {
				{-1, 2}: {},
				{0, 2}: {},
				{1, 2}: {},
			},
			Map {
				{0, 2}: {},
			},
		},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := test.a.Intersect(test.b)
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("unexpected map: want %v, got %v", test.want, got)
			}
		})
	}
}

func TestIntersectLines(t *testing.T) {
	tests := map[int]struct {
		in []Line
		want Map
	}{
		1: {
			[]Line{
				{
					Segment{Point{0, 0}, Point{8, 0}},
					Segment{Point{8, 0}, Point{8, 5}},
					Segment{Point{8, 5}, Point{3, 5}},
					Segment{Point{3, 5}, Point{3, 2}},
				},
				// U7,R6,D4,L4
				{
					Segment{Point{0, 0}, Point{0, 7}},
					Segment{Point{0, 7}, Point{6, 7}},
					Segment{Point{6, 7}, Point{6, 3}},
					Segment{Point{6, 3}, Point{2, 3}},
				},
			},
			Map{
				{0, 0}: {},
				{3, 3}: {},
				{6, 5}: {},
			},
		},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := IntersectLines(test.in...)
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("unexpected map: want %v, got %v", test.want, got)
			}
		})
	}
}
