package sqtri

import (
	"fmt"
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/ebnf"
	"github.com/phyrwork/goadvent/app"
	"io"
	"sort"
	"strconv"
)

var lex = lexer.Must(ebnf.New(`
	Int = digit {digit} .
    Space = " " .
	Newline = "\n" .
	digit = "0"…"9" .
`))

type Triangle struct {
	A int `Space* @Int`
	B int `Space+ @Int`
	C int `Space+ @Int`
}

type triangles struct {
	Triangles []Triangle `(@@ (Newline @@)*)? Newline?`
}

type Triangles []Triangle

func (a Triangles) Filter(f func (Triangle) bool) Triangles {
	o := make([]Triangle, 0)
	for _, t := range a {
		if f(t) {
			o = append(o, t)
		}
	}
	return o
}

var pars = participle.MustBuild(&triangles{}, participle.Lexer(lex))

func Read(r io.Reader) ([]Triangle, error) {
	l := &triangles{}
	if err := pars.Parse(r, l); err != nil {
		return nil, fmt.Errorf("parse error: %v", err)
	}
	return l.Triangles, nil
}

func (t Triangle) Valid() bool {
	a := []int{t.A, t.B, t.C}
	sort.Slice(a, func (i, j int) bool {
		return a[i] < a[j]
	})
	return a[0] + a[1] > a[2]
}

func NewSolver(f func (Triangle) bool) app.SolverFunc {
	return func (r io.Reader) (string, error) {
		a, err := Read(r)
		if err != nil {
			return "", nil
		}
		b := Triangles(a).Filter(f)
		return strconv.Itoa(len(b)), nil
	}
}
