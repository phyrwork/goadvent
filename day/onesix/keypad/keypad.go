package keypad

import (
	"bufio"
	"fmt"
	"github.com/phyrwork/goadvent/app"
	"github.com/phyrwork/goadvent/vector"
	"io"
)

type Position [2]int

func NewPosition(d ...int) Position {
	var c Position
	copy(c[:], d)
	return c
}

func (p Position) Move(q Position) Position {
	return NewPosition(vector.Sum(p[:], q[:])...)
}

type Keypad map[Position]rune

type Cursor struct {
	kp Keypad
	p  Position
}

func NewCursor(kp Keypad, start Position) (*Cursor, error) {
	if _, ok := kp[start]; !ok {
		return nil, fmt.Errorf("no key at %v", start)
	}
	return &Cursor{kp, start}, nil
}

func (cur Cursor) Key() rune { return cur.kp[cur.p] }

func (cur *Cursor) Move(q Position) {
	p := cur.p.Move(q)
	if _, ok := cur.kp[p]; ok {
		cur.p = p
	}
}

type DecoderFn func (*Decoder)

type Grammar map[rune]DecoderFn

const EOF rune = -1

type Decoder struct {
	gr  Grammar
	cur *Cursor
	out []rune
}

func NewDecoder(gr Grammar, kp Keypad, sp Position) *Decoder {
	cur, err := NewCursor(kp, sp)
	if err != nil {
		return nil
	}
	return &Decoder{
		gr:  gr,
		cur: cur,
		out: make([]rune, 0),
	}
}

func (d *Decoder) Emit() { d.out = append(d.out, d.cur.Key()) }

func (d *Decoder) Decode(r io.Reader) error {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanRunes)
	for sc.Scan() {
		c := []rune(sc.Text())[0]
		h := d.gr[c]
		if h == nil {
			return fmt.Errorf("unhandled instruction %v", c)
		}
		h(d)
	}
	if err := sc.Err(); err != nil {
		return fmt.Errorf("scan error: %v", err)
	}
	//if h := d.gr[EOF]; h != nil {
	//	h(d)
	//}
	return nil
}

func (d *Decoder) Out() []rune {
	out := make([]rune, len(d.out))
	copy(out, d.out)
	return out
}

var DefaultGrammar = Grammar{
	'U': func (d *Decoder) { d.cur.Move(Position{0, 1}) },
	'D': func (d *Decoder) { d.cur.Move(Position{0, -1}) },
	'L': func (d *Decoder) { d.cur.Move(Position{-1, 0}) },
	'R': func (d *Decoder) { d.cur.Move(Position{1, 0}) },
	'\n': func (d *Decoder) { d.Emit() },
	EOF: func (d *Decoder) { d.Emit() },
}

var SquareKeypad = Keypad{
	{0, 2}: '1',
	{1, 2}: '2',
	{2, 2}: '3',
	{0, 1}: '4',
	{1, 1}: '5',
	{2, 1}: '6',
	{0, 0}: '7',
	{1, 0}: '8',
	{2, 0}: '9',
}

var DiamondKeypad = Keypad{
	{2, 4}: '1',
	{1, 3}: '2',
	{2, 3}: '3',
	{3, 3}: '4',
	{0, 2}: '5',
	{1, 2}: '6',
	{2, 2}: '7',
	{3, 2}: '8',
	{4, 2}: '9',
	{1, 1}: 'A',
	{2, 1}: 'B',
	{3, 1}: 'C',
	{2, 0}: 'D',
}

func NewSolver(kp Keypad, sp Position) app.SolverFunc {
	return func (r io.Reader) (string, error) {
		dc := NewDecoder(DefaultGrammar, kp, sp)
		if err := dc.Decode(r); err != nil {
			return "", fmt.Errorf("decode error: %v", err)
		}
		out := make([]rune, 0)
		for _, k := range dc.Out() {
			out = append(out, k)
		}
		return string(out), nil
	}
}
