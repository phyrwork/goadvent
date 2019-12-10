package grid2

import "testing"

func TestIntersectSegment2(t *testing.T) {
	tests := map[string]struct {
		a, b Segment
		x    Segment
		ok   bool
	}{
		"diagonal X": {
			Segment{Point{10,0}, Point{0,10}},
			Segment{Point{0,0}, Point{10,10}},
			Segment{Point{5,5}, Point{5,5}},
			true,
		},
		"rectilinear X": {
			Segment{Point{1,0}, Point{1,2}},
			Segment{Point{0,1}, Point{2,1}},
			Segment{Point{1,1}, Point{1,1}},
			true,
		},
		"rectilinear T": {
			Segment{Point{0,5}, Point{5,5}},
			Segment{Point{3,0}, Point{3,5}},
			Segment{Point{3,5}, Point{3,5}},
			true,
		},
		"rectilinear L": {
			Segment{Point{0,5}, Point{0,0}},
			Segment{Point{0,0}, Point{5,0}},
			Segment{Point{0,0}, Point{0,0}},
			true,
		},
		"colinear point horizontal": {
			Segment{Point{0,0}, Point{0,3}},
			Segment{Point{0,3}, Point{0,5}},
			Segment{Point{0,3}, Point{0,3}},
			true,
		},
		"colinear point vertical": {
			Segment{Point{0,0}, Point{3,0}},
			Segment{Point{3,0}, Point{5,0}},
			Segment{Point{3,0}, Point{3,0}},
			true,
		},
		"colinear point diag": {
			Segment{Point{0,0}, Point{3,3}},
			Segment{Point{3,3}, Point{5,5}},
			Segment{Point{3,3}, Point{3,3}},
			true,
		},
		"colinear segment horizontal": {
			Segment{Point{0,0}, Point{2,0}},
			Segment{Point{1,0}, Point{4,0}},
			Segment{Point{1,0}, Point{2,0}},
			true,
		},
		"colinear segment vertical": {
			Segment{Point{0,0}, Point{0,2}},
			Segment{Point{0,1}, Point{0,4}},
			Segment{Point{0,1}, Point{0,2}},
			true,
		},
		"colinear segment diagonal 1": { // 45deg
			Segment{Point{0,0}, Point{2,2}},
			Segment{Point{1,1}, Point{4,4}},
			Segment{Point{1,1}, Point{2,2}},
			true,
		},
		"colinear segment diagonal 2": {
			Segment{Point{0,0}, Point{2,4}},
			Segment{Point{1,2}, Point{3,6}},
			Segment{Point{1,2}, Point{2,4}},
			true,
		},
		"colinear segment horizontal (a contains b)": {
			Segment{Point{0,0}, Point{4,0}},
			Segment{Point{0,0}, Point{2,0}},
			Segment{Point{0,0}, Point{2,0}},
			true,
		},
		"colinear segment vertical (a contains b)": {
			Segment{Point{0,0}, Point{0,4}},
			Segment{Point{0,0}, Point{0,2}},
			Segment{Point{0,0}, Point{0,2}},
			true,
		},
		"colinear segment horizontal (a equals b)": {
			Segment{Point{0,0}, Point{4,0}},
			Segment{Point{0,0}, Point{4,0}},
			Segment{Point{0,0}, Point{4,0}},
			true,
		},
		"colinear segment vertical (a equals b)": {
			Segment{Point{0,0}, Point{0,4}},
			Segment{Point{0,0}, Point{0,4}},
			Segment{Point{0,0}, Point{0,4}},
			true,
		},
		"parallel": {
			Segment{Point{0,0}, Point{0,4}},
			Segment{Point{2,0}, Point{2,4}},
			Segment{},
			false,
		},
		"no intersection rectilinear T": {
			Segment{Point{0,5}, Point{5,5}},
			Segment{Point{3,0}, Point{3,4}},
			Segment{},
			false,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			x, ok := IntersectSegment(test.a, test.b)
			if ok != test.ok {
				t.Fatalf("unexpected intersect detect: want %v, got %v", test.ok, ok)
			}
			if !SegmentEq(test.x, x) {
				t.Fatalf("unexpected intersect segment: want %v, got %v", test.x, x)
			}
		})
	}
}

