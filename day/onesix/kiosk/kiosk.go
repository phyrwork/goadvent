package kiosk

import (
	"fmt"
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/ebnf"
	"io"
)

var lex = lexer.Must(ebnf.New(`
	Sep = "\n" .
	Punct = "-" | "[" | "]" .
	Word = alpha {alpha} .
	Int = digit {digit} .
	alpha = "a"…"z" .
	digit = "0"…"9" .
`))

type Room struct {
	Name   string `@(Word ("-" Word)*)`
	Sector int    `"-" @Int`
	Hash   string `"[" @Word "]"`
}

type rooms struct {
	Rooms []Room `(@@ ("\n" @@)* ("\n")?)?`
}

var parser = participle.MustBuild(&rooms{}, participle.Lexer(lex))

func Read(r io.Reader) ([]Room, error) {
	l := &rooms{}
	if err := parser.Parse(r, l); err != nil {
		return nil, fmt.Errorf("parse error: %v", err)
	}
	// TODO: fix parser, remove workaround
	// for some reason the last hyphen is getting lumped in with
	// the room name... weird.
	for i := range l.Rooms {
		name := l.Rooms[i].Name
		name = name[:len(name)-1]
		l.Rooms[i].Name = name
	}
	return l.Rooms, nil
}
