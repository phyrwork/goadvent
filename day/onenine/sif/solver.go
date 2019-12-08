package sif

import (
	"fmt"
	"github.com/phyrwork/goadvent/app"
	"github.com/phyrwork/goadvent/iterator"
	"io"
)

func Read(s io.Reader, h, w int) (Layers, error) {
	l := make(Layers, 0)
	ok := true
	scr := iterator.NewRuneScanner(s)
	scr.Skip = iterator.SkipWhitespace
	sct := iterator.NewTransformIterator(scr, func (v interface{}) (interface{}, error) {
		c := v.(rune)
		return c - '0', nil
	})
	top:
	for {
		y := make([][]rune, h)
		for r := range y {
			x := make([]rune, w)
			for c := range x {
				if !sct.Next() {
					break top
				}
				ok = false
				x[c] = sct.Value().(rune)
			}
			y[r] = x
		}
		l = append(l, y)
		ok = true
	}
	if err := sct.Err(); err != nil {
		return nil, fmt.Errorf("scan error: %v", err)
	}
	if !ok {
		return nil, fmt.Errorf("incomplete layer")
	}
	return l, nil
}

func solve1(r io.Reader, h, w int) app.Solution {
	l, err := Read(r, h, w)
	if err != nil {
		return app.Errorf("read error: %v", err)
	}
	if len(l) == 0 {
		return app.Errorf("no layers")
	}
	var lm Layer
	lc := 0
	for _, l := range l {
		c := Count(l, func (d rune) bool { return d == 0 })
		if lm == nil {
			lm, lc = l, c
			continue
		}
		if c < lc {
			lm, lc = l, c
		}
	}
	co := Count(lm, func (d rune) bool { return d == 1 })
	ct := Count(lm, func (d rune) bool { return d == 2 })
	return app.Int(co * ct)
}

func Solve1(r io.Reader) app.Solution {
	return solve1(r, 6, 25)
}
