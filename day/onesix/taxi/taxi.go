package taxi

import (
	"fmt"
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/ebnf"
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

func Walk(c Compass, s ...Step) vector.Vector {
	v := vector.Vector{0, 0}
	for _, s := range s {
		c = c.Turn(s.Dir)
		w := vector.Mult(c.Vector(), s.Dist)
		v = vector.Sum(v, w)
	}
	return v
}

func SolveDist(r io.Reader) (string, error) {
	s, err := Read(r)
	if err != nil {
		return "", err
	}
	v := Walk(North, s...)
	d := vector.Manhattan(vector.Vector{0, 0}, v)
	return strconv.Itoa(d), nil
}
