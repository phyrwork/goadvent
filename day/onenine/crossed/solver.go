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

func Solve1(r io.Reader) app.Solution {
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
	// find closest
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

func Solve2(r io.Reader) app.Solution {
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
	if len(m) == 0 {
		return app.Errorf("lines do not intersect")
	}
	// line point distances
	dc := make([]map[Point2]int, len(c))
	for i := range c {
		dc[i] = make(map[Point2]int)
		for d, p := range c[i].(DirectedLine).OrderedPoints() {
			dc[i][p] = d
		}
	}
	// find best combination
	sum := func (p Point2) int {
		s := 0
		for i := range dc {
			s += dc[i][p]
		}
		return s
	}
	s := m.Points().Slice()
	min := sum(s[0])
	for i := 1; i < len(s); i++ {
		cur := sum(s[i])
		if cur < min {
			min = cur
		}
	}
	return app.Int(min)
}