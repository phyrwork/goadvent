package op

import (
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/ebnf"
	"io"
)

var grammar = lexer.Must(ebnf.New(`
	Ident = alpha .
	Integer = [ "-" ] digit { digit } .
	Punct = "-" .
	Whitespace = ( " " | "\t" ) { " " | "\t" } .
	EOL = "\n" .
	alpha = "a"…"z" .
    digit = "0"…"9" .
`))

type ops struct {
	Ops []Op `(@@ (EOL @@)* (EOL)?)?`
}

var parser = participle.MustBuild(&ops{}, participle.Lexer(grammar), participle.Elide("Whitespace"))

func Parse(r io.Reader) ([]Op, error) {
	o := &ops{}
	return o.Ops, parser.Parse(r, o)
}