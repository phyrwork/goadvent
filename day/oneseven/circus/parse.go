package circus

import (
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/ebnf"
	"io"
)

type Descriptor struct {
	Id        string   `@Identifier`
	Weight    int      `"(" @Integer ")"`
	Subtowers []string `("-" ">" @Identifier ("," @Identifier)*)?`
}

type descriptorList struct {
	Descs []Descriptor `(@@ ("\n" @@)* ("\n")?)?`
}

var grammar = lexer.Must(ebnf.New(`
	Identifier = alpha { alpha } .
    Integer = digit { digit } .
	Punct = "(" | ")" | "-" | ">" | "," | "\n" .
	Whitespace = " " .
    alpha = "a"…"z" | "A"…"Z" .
    digit = "0"…"9" .

`))

var parser = participle.MustBuild(&descriptorList{}, participle.Lexer(grammar), participle.Elide("Whitespace"))

func Parse(r io.Reader) ([]Descriptor, error) {
	l := &descriptorList{}
	return l.Descs, parser.Parse(r, l)
}
