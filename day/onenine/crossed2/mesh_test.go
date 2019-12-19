package crossed2

import (
	"reflect"
	"strconv"
	"testing"
)

func TestVectorLineBuilder(t *testing.T) {
	tests := map[int]struct {
		in   []Vector
		from Point
		want Wire
	}{
		// R8,U5,L5,D3
		1: {
			[]Vector{{8,0},{0,5},{-5,0},{0,-3}},
			Point{0,0},
			Wire{
				{Point{0,0},Point{8,0}}:{},
				{Point{8,0},Point{8,5}}:{},
				{Point{8,5},Point{3,5}}:{},
				{Point{3,5},Point{3,2}}:{},
			},
		},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := NewWire(test.from, test.in...)
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("unexpected Segments: want %v, got %v", test.want, got)
			}
		})
	}
}

func TestWire_Intersect(t *testing.T) {
	tests := map[int]struct {
		a, b Wire
		want map[Point]struct{}
	}{
		1: {
			NewWire(Point{0, 0}, []Vector{{8, 0}, {0, 5}, {-5, 0}, {0, -3}}...),
			NewWire(Point{0, 0}, []Vector{{0, 7}, {6, 0}, {0, -4}, {-4, 0}}...),
			map[Point]struct{}{
				{0, 0}: {},
				{3, 3}: {},
				{6, 5}: {},
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

func TestNewMesh(t *testing.T) {
	tests := map[int]struct {
		wires    []Wire
		points   map[Point]struct{}
		segments map[Segment]struct{}
	}{
		1: {
			[]Wire{
				NewWire(Point{0, 0}, []Vector{{8, 0}, {0, 5}, {-5, 0}, {0, -3}}...),
				NewWire(Point{0, 0}, []Vector{{0, 7}, {6, 0}, {0, -4}, {-4, 0}}...),
			},
			map[Point]struct{}{
				// a
				Point{0,0}:{},
				Point{8,0}:{},
				Point{8,0}:{},
				Point{8,5}:{},
				Point{8,5}:{},
				Point{3,5}:{},
				Point{3,5}:{},
				Point{3,2}:{},
				// b
				Point{0,7}:{},
				Point{6,7}:{},
				Point{6,3}:{},
				Point{2,3}:{},
				// x
				Point{3,3}:{},
				Point{6,5}:{},
			},
			map[Segment]struct{}{
				{Point{0,0},Point{8,0}}:{},
				{Point{8,0},Point{8,5}}:{},
				{Point{8,5},Point{6,5}}:{},
				{Point{6,5},Point{3,5}}:{},
				{Point{3,5},Point{3,3}}:{},
				{Point{3,3},Point{3,2}}:{},
				{Point{0,0},Point{0,7}}:{},
				{Point{0,7},Point{6,7}}:{},
				{Point{6,7},Point{6,5}}:{},
				{Point{6,5},Point{6,3}}:{},
				{Point{6,3},Point{3,3}}:{},
				{Point{3,3},Point{2,3}}:{},
			},
		},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			m := NewMesh(test.wires...)
			points := m.Points()
			if !reflect.DeepEqual(test.points, points) {
				t.Fatalf("unexpected points: want %v, got %v", test.points, points)
			}
			segments := m.Segments()
			if !Segments(segments).Eq(test.segments) {
				t.Fatalf("unexpected segments: want %v, got %v", test.segments, segments)
			}
		})
	}
}
