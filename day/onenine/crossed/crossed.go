package crossed

import (
	"bufio"
	"fmt"
	"github.com/phyrwork/goadvent/app"
	"github.com/phyrwork/goadvent/vector"
	"io"
	"regexp"
	"strconv"
	"strings"
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

func (s Segment) Segment() Segment { return s }

func (s Segment) Points() []Point {
	r := vector.NewRange(Vector(s.S), Vector(s.E))
	l := make([]Point, 0)
	if err := r.Each(func (v Vector) error {
		p := Point(v)
		l = append(l, p)
		return nil
	}); err != nil {
		// nil error func given
		panic("each point error")
	}
	return l
}

type PointVector struct {
	P Point
	V Vector
}

func (p PointVector) To() Point {
	return Point(vector.Sum(p.V, Vector(p.P)))
}

func (p PointVector) Segment() Segment {
	return Segment {p.P, p.To()}
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

type Map map[Point2]struct{}

func (m Map) Add(l ...Point) {
	for _, p := range l {
		m[NewPoint2(p...)] = struct{}{}
	}
}

func (m Map) Intersect(n Map) Map {
	if m == nil {
		return nil
	}
	o := make(Map)
	for p := range m {
		if _, ok := n[p]; ok {
			o[p] = struct{}{}
		}
	}
	return o
}

func (m Map) Filter(f func (Point2) bool) Map {
	o := make(Map)
	for p := range m {
		if f(p) {
			o[p] = struct{}{}
		}
	}
	return o
}

var vectorRegex = regexp.MustCompile(`(U|D|L|R)(\d+)`)

func ReadVector(s string) (Vector, error) {
	m := vectorRegex.FindStringSubmatch(s)
	if len(m) < 3 {
		return Vector{}, fmt.Errorf("not a point vector string: %v", s)
	}
	d, err := strconv.Atoi(m[2])
	if err != nil {
		return Vector{}, fmt.Errorf("atoi error: %v", err)
	}
	var v Vector
	switch m[1] {
	case "U":
		v = Vector{0, 1}
	case "D":
		v = Vector{0, -1}
	case "L":
		v = Vector{-1, 0}
	case "R":
		v = Vector{1, 0}
	default:
		return Vector{}, fmt.Errorf("unknown direction: %v", m[1])
	}
	return vector.Mult(v, d), nil
}

func Read(r io.Reader) ([][]Vector, error) {
	o := make([][]Vector, 0)
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		s := sc.Text()
		w := strings.Split(s, ",")
		l := make([]Vector, len(w))
		for i, w := range w {
			v, err := ReadVector(w)
			if err != nil {
				return nil, fmt.Errorf("vector read error: %v", err)
			}
			l[i] = v
		}
		o = append(o, l)
	}
	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("scan error: %v", err)
	}
	return o, nil
}

type Segmenter interface {
	Segment() Segment
}

func Segments(s ...Segmenter) []Segment {
	o := make([]Segment, len(s))
	for i, s := range s {
		o[i] = s.Segment()
	}
	return o
}

type Line []Segmenter

func (l Line) Map() Map {
	m := make(Map)
	for _, s := range l {
		m.Add(s.Segment().Points()...)
	}
	return m
}

type VectorLineBuilder struct {
	From Point
}

func (b VectorLineBuilder) New(v ...Vector) Line {
	l := make(Line, len(v))
	cur := b.From
	for i, v := range v {
		s := PointVector{cur, v}
		l[i] = s
		cur = s.To()
	}
	return l
}

func IntersectLines(l ...Line) Map {
	var m Map
	if len(l) >= 1 {
		l := l[0]
		m = l.Map()
	}
	for i := 1; i < len(l); i++ {
		l := l[i]
		m = m.Intersect(l.Map())
	}
	return m
}

func Solve(r io.Reader) app.Solution {
	// line vectors
	v, err := Read(r)
	if err != nil {
		return app.Errorf("input read error: %v", err)
	}
	// make lines
	o := Point{0, 0}
	c := make([]Line, len(v))
	for i := range v {
		b := VectorLineBuilder{o}
		l := b.New(v[i]...)
		c[i] = l
	}
	// intersect lines
	m := IntersectLines(c...).Filter(func (p Point2) bool {
		return !vector.Eq(Vector(NewPoint(p[:]...)), Vector(o))
	})
	var p Point
	var dp int
	for q := range m {
		q := NewPoint(q[:]...)
		// best by default
		if p == nil {
			p = q
			dp = vector.Manhattan(Vector(o), Vector(p))
			continue
		}
		// best by distance
		if dq := vector.Manhattan(Vector(o), Vector(q)); dq < dp {
			p = q
			dp = dq
		}
	}
	if p == nil {
		return app.Errorf("lines do not intersect")
	}
	return app.Int(dp)
}