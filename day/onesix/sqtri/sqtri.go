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

type Triangle struct {
	A int
	B int
	C int
}

type Reader func (io.Reader) (Triangles, error)

var lex = lexer.Must(ebnf.New(`
	Int = digit {digit} .
    Space = " " .
	Newline = "\n" .
	digit = "0"…"9" .
`))

type Row struct {
	A int `Space* @Int`
	B int `Space+ @Int`
	C int `Space+ @Int`
}

type rows struct {
	Rows []Row `(@@ (Newline @@)*)? Newline?`
}

var parsr = participle.MustBuild(&rows{}, participle.Lexer(lex))

func ReadRows(r io.Reader) (Triangles, error) {
	l := &rows{}
	if err := parsr.Parse(r, l); err != nil {
		return nil, fmt.Errorf("parse error: %v", err)
	}
	a := make(Triangles, len(l.Rows))
	for i, r := range l.Rows {
		a[i] = Triangle{r.A, r.B, r.C}
	}
	return a, nil
}

type Trio struct {
	R Row `        @@`
	S Row `Newline @@`
	T Row `Newline @@`
}

type trios struct {
	Trios []Trio `(@@ (Newline @@)*)? Newline?`
}

var parst = participle.MustBuild(&trios{}, participle.Lexer(lex))

func ReadTrios(r io.Reader) (Triangles, error) {
	l := &trios{}
	if err := parst.Parse(r, l); err != nil {
		return nil, fmt.Errorf("parse error: %v", err)
	}
	a := make(Triangles, 0, len(l.Trios) * 3)
	for _, t := range l.Trios {
		a = append(a, Triangle{t.R.A, t.S.A, t.T.A})
		a = append(a, Triangle{t.R.B, t.S.B, t.T.B})
		a = append(a, Triangle{t.R.C, t.S.C, t.T.C})
	}
	return a, nil
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

func (t Triangle) Valid() bool {
	a := []int{t.A, t.B, t.C}
	sort.Slice(a, func (i, j int) bool {
		return a[i] < a[j]
	})
	return a[0] + a[1] > a[2]
}

func NewSolver(f Reader) app.SolverFunc {
	return func (r io.Reader) (string, error) {
		a, err := ReadRows(r)
		if err != nil {
			return "", nil
		}
		b := Triangles(a).Filter(Triangle.Valid)
		return strconv.Itoa(len(b)), nil
	}
}
