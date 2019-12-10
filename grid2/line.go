package grid2

func orientation(p, q, r Point) int {
	// gradient of segment (p, q): σ = (y2 - y1)/(x2 - x1)
	// gradient of segment (q, r): τ = (y3 - y2)/(x3 - x2)
	// σ > τ: clockwise
	// σ > τ: counter-clockwise
	// σ = τ: co-linear
	// https://www.geeksforgeeks.org/check-if-two-given-line-segments-intersect/
	//
	// above re-arranges into the following:
	px, py := p.XY()
	qx, qy := q.XY()
	rx, ry := r.XY()
	o := (qy-py)*(rx-qx)-(qx-px)*(ry-qy)
	if o > 0 {
		return Clockwise
	} else if o < 0 {
		return CounterClockwise
	} else {
		return Colinear
	}
}

func colinearContains (p, q, r Point) bool {
	// for co-linear points (p, q, r)
	// check q is on segment (p, r)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	px, py := p.XY()
	qx, qy := q.XY()
	rx, ry := r.XY()
	if qx > max(px, rx) {
		return false
	}
	if qx < min(px, rx) {
		return false
	}
	if qy > max(py, ry) {
		return false
	}
	if qy < min(py, ry) {
		return false
	}
	return true
}

// Point is to Vector as Time is to Interval
type Point [2]int

func (a Point) XY() (int, int) {
	return a[0], a[1]
}

func (a Point) Sum(b Vector) Point {
	return Point{
		a[0] + b[0],
		a[1] + b[1],
	}
}

func (a Point) Sub(b Point) Point {
	return Point{a[0]-b[0], a[1]-b[1]}
}

type Segment struct {
	Start Point
	End   Point
}

func (s Segment) Ray() Ray {
	return Ray{
		s.Start,
		Vector(s.End).Sub(Vector(s.Start)),
	}
}

func (s Segment) Point() (Point, bool) {
	if s.Start == s.End {
		return s.Start, true
	}
	return Point{}, false
}

func (s Segment) Reverse() Segment {
	return Segment{s.End, s.Start}
}

func (s Segment) Contains(p Point) bool {
	return orientation(s.Start, p, s.End) == Colinear && colinearContains(s.Start, p, s.End)
}

func (s Segment) Intersect(t Segment) (Segment, bool) {
	return IntersectSegment(s, t)
}

func (s Segment) Eq(t Segment) bool {
	return SegmentEq(s, t)
}

func SegmentEq(a, b Segment) bool {
	if a.Start == b.Start && a.End == b.End {
		return true
	}
	if a.Start == b.End && a.End == b.Start {
		return true
	}
	return false
}

type Ray struct {
	Start Point
	Dir   Vector
}

func (r Ray) Segment(d int) Segment {
	return Segment{
		r.Start,
		Point(r.Dir.Mul(d)),
	}
}

// A Line is a Ray that extends in both directions
type Line Ray

func intersectLine(p1, q1, p2, q2 Point) (Point, bool) {
	// line points
	x1, y1 := p1.XY()
	x2, y2 := q1.XY()
	x3, y3 := p2.XY()
	x4, y4 := q2.XY()
	// intersect
	// https://en.wikipedia.org/wiki/Line%E2%80%93line_intersection#Given_two_points_on_each_line
	d := (x1-x2)*(y3-y4)-(y1-y2)*(x3-x4)
	if d == 0 {
		return Point{}, false
	}
	n12 := (x1*y2)-(y1*x2)
	n34 := (x3*y4)-(y3*x4)
	npx := (n12*(x3-x4))-((x1-x2)*n34)
	npy := (n12*(y3-y4))-((y1-y2)*n34)
	px := npx/d
	py := npy/d
	return Point{px, py}, true
}

func IntersectLine(a, b Line) (Point, bool) {
	return intersectLine(a.Start, Point(a.Dir), b.Start, Point(b.Dir))
}

const (
	Clockwise = 1
	CounterClockwise = -1
	Colinear = 0
)

func IntersectSegment(a, b Segment) (Segment, bool) {
	p1, q1 := a.Start, a.End
	p2, q2 := b.Start, b.End
	// four orientations required to solve
	// general and co-linear cases
	o1 := orientation(p1, q1, p2)
	o2 := orientation(p1, q1, q2)
	o3 := orientation(p2, q2, p1)
	o4 := orientation(p2, q2, q1)
	if o1 != o2 && o3 != o4 {
		// point intersection
		p, _ := intersectLine(a.Start, a.End, b.Start, b.End)
		return Segment{p, p}, true
	}
	// collect unique intersection points
	// of co-linear segments
	// https://www.geeksforgeeks.org/check-if-two-given-line-segments-intersect/
	if o1 != Colinear && o2 != Colinear && o3 != Colinear && o4 != Colinear {
		return Segment{}, false
	}
	xm := make(map[Point]struct{}, 4)
	if o1 == Colinear && colinearContains(p1, p2, q1) {
		// p2 lies on p1q1
		xm[p2] = struct{}{}
	}
	if o2 == Colinear && colinearContains(p1, q2, q1) {
		// q1 lies on p1q1
		xm[q2] = struct{}{}
	}
	if o3 == Colinear && colinearContains(p2, p1, q2) {
		// p1 lies on p2q2
		xm[p1] = struct{}{}
	}
	if o4 == Colinear && colinearContains(p2, q1, q2) {
		// q1 lies p2q1
		xm[q1] = struct{}{}
	}
	xn := make([]Point, 0, len(xm))
	for x := range xm {
		xn = append(xn, x)
	}
	switch len(xn) {
	case 0:
		// no intersection
		return Segment{}, false
	case 1:
		// co-linear point intersection
		return Segment{xn[0], xn[0]}, true
	case 2:
		// co-linear segment intersection
		return Segment{xn[0], xn[1]}, true
	default:
		// more than two points is A Bad Thing™ and
		// means something probably went wrong elsewhere...
		panic(`segments intersect at more than two points'`)
	}
}
