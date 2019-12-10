package grid2

import "testing"

func TestIntersectSegment2(t *testing.T) {
	tests := map[string]struct {
		a, b Segment2
		x    Segment2
		ok   bool
	}{
		"diagonal X": {
			Segment2{Point{10,0}, Point{0,10}},
			Segment2{Point{0,0}, Point{10,10}},
			Segment2{Point{5,5}, Point{5,5}},
			true,
		},
		"rectilinear X": {
			Segment2{Point{1,0}, Point{1,2}},
			Segment2{Point{0,1}, Point{2,1}},
			Segment2{Point{1,1}, Point{1,1}},
			true,
		},
		"rectilinear T": {
			Segment2{Point{0,5}, Point{5,5}},
			Segment2{Point{3,0}, Point{3,5}},
			Segment2{Point{3,5}, Point{3,5}},
			true,
		},
		"rectilinear L": {
			Segment2{Point{0,5}, Point{0,0}},
			Segment2{Point{0,0}, Point{5,0}},
			Segment2{Point{0,0}, Point{0,0}},
			true,
		},
		"colinear point horizontal": {
			Segment2{Point{0,0}, Point{0,3}},
			Segment2{Point{0,3}, Point{0,5}},
			Segment2{Point{0,3}, Point{0,3}},
			true,
		},
		"colinear point vertical": {
			Segment2{Point{0,0}, Point{3,0}},
			Segment2{Point{3,0}, Point{5,0}},
			Segment2{Point{3,0}, Point{3,0}},
			true,
		},
		"colinear point diag": {
			Segment2{Point{0,0}, Point{3,3}},
			Segment2{Point{3,3}, Point{5,5}},
			Segment2{Point{3,3}, Point{3,3}},
			true,
		},
		"colinear segment horizontal": {
			Segment2{Point{0,0}, Point{2,0}},
			Segment2{Point{1,0}, Point{4,0}},
			Segment2{Point{1,0}, Point{2,0}},
			true,
		},
		"colinear segment vertical": {
			Segment2{Point{0,0}, Point{0,2}},
			Segment2{Point{0,1}, Point{0,4}},
			Segment2{Point{0,1}, Point{0,2}},
			true,
		},
		"colinear segment diagonal 1": { // 45deg
			Segment2{Point{0,0}, Point{2,2}},
			Segment2{Point{1,1}, Point{4,4}},
			Segment2{Point{1,1}, Point{2,2}},
			true,
		},
		"colinear segment diagonal 2": {
			Segment2{Point{0,0}, Point{2,4}},
			Segment2{Point{1,2}, Point{3,6}},
			Segment2{Point{1,2}, Point{2,4}},
			true,
		},
		"colinear segment horizontal (a contains b)": {
			Segment2{Point{0,0}, Point{4,0}},
			Segment2{Point{0,0}, Point{2,0}},
			Segment2{Point{0,0}, Point{2,0}},
			true,
		},
		"colinear segment vertical (a contains b)": {
			Segment2{Point{0,0}, Point{0,4}},
			Segment2{Point{0,0}, Point{0,2}},
			Segment2{Point{0,0}, Point{0,2}},
			true,
		},
		"colinear segment horizontal (a equals b)": {
			Segment2{Point{0,0}, Point{4,0}},
			Segment2{Point{0,0}, Point{4,0}},
			Segment2{Point{0,0}, Point{4,0}},
			true,
		},
		"colinear segment vertical (a equals b)": {
			Segment2{Point{0,0}, Point{0,4}},
			Segment2{Point{0,0}, Point{0,4}},
			Segment2{Point{0,0}, Point{0,4}},
			true,
		},
		"parallel": {
			Segment2{Point{0,0}, Point{0,4}},
			Segment2{Point{2,0}, Point{2,4}},
			Segment2{},
			false,
		},
		"no intersection rectilinear T": {
			Segment2{Point{0,5}, Point{5,5}},
			Segment2{Point{3,0}, Point{3,4}},
			Segment2{},
			false,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			x, ok := IntersectSegment2(test.a, test.b)
			if ok != test.ok {
				t.Fatalf("unexpected intersect detect: want %v, got %v", test.ok, ok)
			}
			if !Segment2Eq(test.x, x) {
				t.Fatalf("unexpected intersect segment: want %v, got %v", test.x, x)
			}
		})
	}
}

