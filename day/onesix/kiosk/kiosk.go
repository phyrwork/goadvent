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

// TODO: for some reason this grammar is including the hyphen
//  separating the room name from the sector with the room name...
//  fixme!
type room struct {
	Name   string `@(Word ("-" Word)*)`
	Sector int    `"-" @Int`
	Hash   string `"[" @Word "]"`
}

type rooms struct {
	Rooms []room `(@@ ("\n" @@)* ("\n")?)?`
}

type Room struct {
	Name   string
	Sector int
	Hash   Hash // Not sure we can output to this directly from parser
}

var parser = participle.MustBuild(&rooms{}, participle.Lexer(lex))

func Read(r io.Reader) ([]Room, error) {
	l := &rooms{}
	if err := parser.Parse(r, l); err != nil {
		return nil, fmt.Errorf("parse error: %v", err)
	}
	// TODO: fix parser, remove workaround
	for i := range l.Rooms {
		name := l.Rooms[i].Name
		name = name[:len(name)-1]
		l.Rooms[i].Name = name
	}
	out := make([]Room, len(l.Rooms))
	for i, r := range l.Rooms {
		chk, err := StringToHash(r.Hash)
		if err != nil {
			return nil, fmt.Errorf("format error: %v", err)
		}
		out[i] = Room{r.Name, r.Sector, chk}
	}
	return out, nil
}

type Validator HashFn

func (v Validator) Valid(r Room) (bool, error) {
	chk, err := v(r.Name)
	if err != nil {
		switch err := err.(type) {
		case UnhashableError:
			return false, nil
		default:
			// it's a real error
			return false, err
		}
	}
	return chk == r.Hash, nil
}

type FilterFn func (Room) (bool, error)

func Filter(l []Room, f FilterFn) ([]Room, error) {
	o := make([]Room, 0)
	for _, r := range l {
		if ok, err := f(r); err != nil {
			return nil, fmt.Errorf("filter error: %v", err)
		} else if ok {
			o = append(o, r)
		}
	}
	return o, nil
}

type MaskFn func(i int, c rune) bool

type ReplaceFn func (int, rune) rune

func ReplaceRunes(w []rune, mask MaskFn, with func (int, rune) rune) {
	for i, c := range w {
		if mask(i, c) {
			w[i] = with(i, c)
		}
	}
}

func Decrypt(r Room) Room {
	// replace hyphens with spaces
	mask := func (_ int, c rune) bool { return c == '-' }
	with := func (_ int, _ rune) rune { return ' ' }
	// rotate
	ws := WrapShifter{r.Sector, 'a', 'z'}
	ms := MaskShifter{ws, mask}
	rot := ms.Shift([]rune(r.Name))
	// replace
	ReplaceRunes(rot, mask, with)
	r.Name = string(rot)
	return r
}