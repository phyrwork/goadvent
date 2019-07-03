package stream

import (
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/ebnf"
	"github.com/phyrwork/goadvent/app"
	"io"
	"strconv"
)

type Group struct {
	Groups []*Groupable `"{" (@@ ("," @@)*)? "}"`
}

type Groupable struct {
	Group *Group `  @@`
	Trash *Trash `| @@`
}

type Trash struct {
	Trash []*Trashable `"<" (@@ (@@)*)? ">"`
}

// TODO: it works but it needs tidying up
type Trashable struct {
	Chars  string `  @Char`
	Ignore string `| @(Cancel (Char | Cancel | ">"))`
}

var lexr = lexer.Must(ebnf.New(`
	Cancel = "!" .
	Char = "a" | "b" | "e" | "i" | "o" | "u" | "\"" | "," | "'" | "<" | "{" | "}" .
    Punct = ">" .
	Whitespace = "\n" .
`))

var parser = participle.MustBuild(&Group{}, participle.Lexer(lexr), participle.Elide("Whitespace"))

func (g *Group) Count() int {
	n := 1
	for _, g := range g.Groups {
		if g.Group != nil {
			n += g.Group.Count()
		}
	}
	return n
}

func (g *Group) Score(d int) int {
	s := d
	for _, g := range g.Groups {
		if g.Group != nil {
			s += g.Group.Score(d + 1)
		}
	}
	return s
}

func NewSolver(f func (g *Group) int) app.SolverFunc {
	return app.SolverFunc(func (r io.Reader) (string, error) {
		g := &Group{}
		err := parser.Parse(r, g)
		if err != nil {
			return "", err
		}
		i := f(g)
		return strconv.Itoa(i), nil
	})
}