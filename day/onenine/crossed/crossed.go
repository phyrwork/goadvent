package crossed

import (
	"github.com/phyrwork/goadvent/vector"
	"log"
)

const (
	X = 0
	Y = 1
)

type Vector = vector.Vector

func NewVector(d ...int) Vector { return d }

type Point Vector

func NewPoint(d ...int) Point { return d }

type Segment struct { S, E Point }

func (s Segment) Points() Points {
	r := vector.NewRange(Vector(s.S), Vector(s.E))
	m := make(Points)
	if err := r.Each(func (v Vector) error {
		p := NewPoint2(v[:]...)
		m[p] = struct{}{}
		return nil
	}); err != nil {
		// nil error func given
		panic("each point error")
	}
	return m
}

type OrderedSegment []Segment

func (o OrderedSegment) Points() Points {
	m := make(Points)
	for _, s := range o {
		for p := range s.Points() {
			m[p] = struct{}{}
		}
	}
	return m
}

type PointVector struct {
	P Point
	V Vector
}

func (p PointVector) To() Point {
	return Point(vector.Sum(p.V, Vector(p.P)))
}

func (p PointVector) OrderedPoints() OrderedPoints {
	// count non-zero dims, because we can only deal with one
	var dim int
	nzd := 0
	for i := range p.V {
		if p.V[i] != 0 {
			dim = i
			nzd++
		}
	}
	switch nzd {
	case 0:
		// nil vector means it's just a point
		return OrderedPoints{NewPoint2(p.P[:]...)}
	case 1:
		break
	default:
		log.Panicf("point vector not rectilinear: %v", p)
	}
	// generate points
	rem := p.V[dim]
	var sgn int
	if rem < 0 {
		rem = -rem
		sgn = -1
	} else {
		sgn = 1
	}
	step := make(Vector, len(p.V))
	step[dim] = sgn
	cur := NewPoint(p.P[:]...)
	out := make(OrderedPoints, 0, 1 + rem)
	out = append(out, NewPoint2(cur[:]...))
	for rem != 0 {
		cur = Point(vector.Sum(Vector(cur), step))
		out = append(out, NewPoint2(cur[:]...))
		if rem > 0 {
			rem--
		} else {
			rem++
		}
	}
	return out
}

func (p PointVector) Points() Points {
	m := make(Points)
	for _, p := range p.OrderedPoints() {
		m[p] = struct{}{}
	}
	return m
}

type Point2 [2]int

func NewPoint2(d ...int) Point2 {
	x := 0
	if len(d) >= 1 {
		x = d[0]
	}
	y := 0
	if len(d) >= 2 {
		y = d[1]
	}
	return Point2{x, y}
}

func (p Point2) Point() Point { return NewPoint(p[:]...) }

type Points map[Point2]struct{}

func (m Points) Points() Points { return m }

func (m Points) Slice() []Point2 {
	o := make([]Point2, 0, len(m))
	for p := range m {
		o = append(o, p)
	}
	return o
}

func (m Points) Intersect(n Points) Points {
	if m == nil {
		return nil
	}
	o := make(Points)
	for p := range m {
		if _, ok := n[p]; ok {
			o[p] = struct{}{}
		}
	}
	return o
}

func (m Points) Filter(f func (Point2) bool) Points {
	o := make(Points)
	for p := range m {
		if f(p) {
			o[p] = struct{}{}
		}
	}
	return o
}

type OrderedPoints []Point2

func (o OrderedPoints) OrderedPoints() OrderedPoints { return o }

func (o OrderedPoints) Points() Points {
	m := make(Points)
	for _, p := range o {
		m[p] = struct{}{}
	}
	return m
}

type Line interface {
	Points() Points
}

type DirectedLine interface {
	Line
	OrderedPoints() OrderedPoints
}

type OrderedPointVectors []PointVector

func (o OrderedPointVectors) Points() Points {
	m := make(Points, len(o))
	for _, v := range o {
		for _, p := range v.OrderedPoints() {
			m[p] = struct{}{}
		}
	}
	return m
}

func (o OrderedPointVectors) OrderedPoints() OrderedPoints {
	s := make(OrderedPoints, 0)
	if len(o) == 0 {
		return nil
	}
	for _, p := range o[0].OrderedPoints() {
		s = append(s, p)
	}
	for i := 1; i < len(o); i++ {
		o := o[i].OrderedPoints()
		for i := 1; i < len(o); i++ {
			s = append(s, o[i])
		}
	}
	return s
}

type VectorLineBuilder struct {
	From Point
}

func (b VectorLineBuilder) New(v ...Vector) OrderedPointVectors {
	l := make(OrderedPointVectors, len(v))
	cur := b.From
	for i, v := range v {
		s := PointVector{cur, v}
		l[i] = s
		cur = s.To()
	}
	return l
}

func IntersectLines(l ...Line) Points {
	var m Points
	if len(l) >= 1 {
		l := l[0]
		m = l.Points()
	}
	for i := 1; i < len(l); i++ {
		l := l[i]
		m = m.Intersect(l.Points())
	}
	return m
}
