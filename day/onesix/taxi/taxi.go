package taxi

import (
	"fmt"
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/ebnf"
	"github.com/phyrwork/goadvent/app"
	"github.com/phyrwork/goadvent/vector"
	"io"
	"strconv"
)

var lex = lexer.Must(ebnf.New(`
	Turn = "L" | "R" .
	Sep = "," .
	Int = (digit) {digit} .
	Whitespace = " " | "\n" .
	digit = "0"…"9" .
`))

const (
	Left = "L"
	Right = "R"
)

type Step struct {
	Dir  string `@Turn`
	Dist int    `@Int`
}

type Directions struct {
	Steps []Step `(@@ ("," @@)+)?`
}

var parser = participle.MustBuild(&Directions{},
	participle.Lexer(lex),
	participle.Elide("Whitespace"),
)

func Read(r io.Reader) ([]Step, error) {
	a := &Directions{}
	if err := parser.Parse(r, a); err != nil {
		return nil, fmt.Errorf("parse error: %v", err)
	}
	return a.Steps, nil
}

const (
	North = 0
	East = 1
	South = 2
	West = 3
)

type Compass int

func (d Compass) Turn(s string) Compass {
	switch s {
	case Right:
		d++
	case Left:
		d--
	default:
		panic("invalid turn")
	}
	if d < 0 {
		d += 4
	}
	return d % 4
}

func (d Compass) Vector() vector.Vector {
	switch d {
	case North:
		return vector.Vector{0, 1}
	case South:
		return vector.Vector{0, -1}
	case East:
		return vector.Vector{1, 0}
	case West:
		return vector.Vector{-1, 0}
	default:
		panic("invalid compass")
	}
}

func WalkUntil(u func (vector.Vector) bool, c Compass, s ...Step) (vector.Vector, bool) {
	v := vector.Vector{0, 0}
	if u != nil && u(v) {
		return v, true
	}
	for _, s := range s {
		c = c.Turn(s.Dir)
		for n := 0; n < s.Dist; n++ {
			v = vector.Sum(v, c.Vector())
			if u != nil && u(v) {
				return v, true
			}
		}
	}
	return v, false
}

func Walk(c Compass, s ...Step) vector.Vector {
	v, _ := WalkUntil(nil, c, s...)
	return v
}

func WalkUntilRevisit(c Compass, s ...Step) vector.Vector {
	m := map[[2]int]struct{}{}
	v, r := WalkUntil(func (v vector.Vector) bool {
		w := [2]int{v[0], v[1]}
		if _, ok := m[w]; ok {
			return true
		}
		m[w] = struct{}{}
		return false
	}, c, s...)
	if r {
		return v
	}
	return nil
}

func Solve(r io.Reader, f func (c Compass, s ...Step) vector.Vector) (int, error) {
	s, err := Read(r)
	if err != nil {
		return 0, err
	}
	v := f(North, s...)
	if v == nil {
		return 0, fmt.Errorf("nil vector")
	}
	d := vector.Manhattan(vector.Vector{0, 0}, v)
	return d, nil
}

func NewSolver(f func (c Compass, s ...Step) vector.Vector) app.SolverFunc {
	return func (r io.Reader) (string, error) {
		d, err := Solve(r, f)
		if err != nil {
			return "", err
		}
		return strconv.Itoa(d), nil
	}
}
