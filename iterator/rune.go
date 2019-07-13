package iterator

import (
	"bufio"
	"io"
)

type RuneIterator interface {
	Iterator
	Rune() rune
}

func RuneSlice(it RuneIterator) ([]rune, error) {
	a := make([]rune, 0)
	for it.Next() {
		a = append(a, it.Rune())
	}
	return a, it.Err()
}

var SkipWhitespace = map[rune]struct{} {
	'\n': {},
	' ':  {},
	'\r': {},
	'\t': {},
}

type RuneScanner struct {
	sc *bufio.Scanner
	c rune
	err error
	Skip map[rune]struct{}
}

func NewRuneScanner(r io.Reader) *RuneScanner {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanRunes)
	return &RuneScanner{sc, 0, nil, nil}
}

func (it *RuneScanner) Next() bool {
	for it.sc.Scan() {
		c := rune(it.sc.Text()[0])
		if _, ok := it.Skip[c]; !ok {
			it.c = c
			return true
		}
	}
	if err := it.sc.Err(); err != nil {
		it.err = err
	}
	return false
}

func (it *RuneScanner) Value() interface{} { return it.c }

func (it *RuneScanner) Rune() rune { return it.c }

func (it *RuneScanner) Err() error { return it.err }