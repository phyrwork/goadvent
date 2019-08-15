package registers

import (
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/ebnf"
	"io"
)

var grammar = lexer.Must(ebnf.New(`
	Identifier = alpha { alpha } .
    Integer = [ "-" ] digit { digit } .
	Punct = "-" | ">" | "<" | "=" | "!" | "\n" .
	Whitespace = " " .
    alpha = "a"…"z" | "A"…"Z" .
    digit = "0"…"9" .
`))

const (
	OpInc = "inc"
	OpDec = "dec"
)

type OpMod struct {
	Reg string `@Identifier`
	Op  string `@("inc" | "dec")`
	Arg int    `@Integer`
}

const (
	OpEq = "=="
	OpNeq = "!="
	OpGt = ">"
	OpGte = ">="
	OpLt = "<"
	OpLte = "<="
)

type OpCmp struct {
	Reg string `@Identifier`
	Op  string `@("!" "=" | "=" "=" | ">" "=" | "<" "=" | ">" | "<")`
	Arg int    `@Integer`
}

type Stmt struct {
	Op   OpMod `@@`
	Cond OpCmp `"if" @@`
}

type stmts struct {
	Stmts []Stmt `(@@ ("\n" @@)* ("\n")?)?`
}

var parser = participle.MustBuild(&stmts{}, participle.Lexer(grammar), participle.Elide("Whitespace"))

func Parse(r io.Reader) ([]Stmt, error) {
	p := &stmts{}
	return p.Stmts, parser.Parse(r, p)
}