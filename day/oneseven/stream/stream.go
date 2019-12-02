package stream

import (
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/ebnf"
	"github.com/phyrwork/goadvent/app"
	"io"
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

// Count groups
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

func (g *Group) Chars() int {
	n := 0
	for _, g := range g.Groups {
		switch {
		case g.Group != nil:
			n += g.Group.Chars()
		case g.Trash != nil:
			for _, t := range g.Trash.Trash {
				n += len(t.Chars)
			}
		}
	}
	return n
}

func NewSolver(f func (g *Group) int) app.SolverFunc {
	return app.SolverFunc(func (r io.Reader) app.Solution {
		g := &Group{}
		err := parser.Parse(r, g)
		if err != nil {
			return app.NewError(err)
		}
		i := f(g)
		return app.Int(i)
	})
}