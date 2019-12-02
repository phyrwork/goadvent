package noise

import (
	"fmt"
	"github.com/phyrwork/goadvent/app"
	"io"
)

func newColumnMap(r Reader) ([]map[rune]int, error) {
	var m []map[rune]int
	add := func (s string) {
		for i, r := range s {
			m[i][r] = m[i][r] + 1
		}
	}
	if r.Next() {
		s := r.String()
		m = make([]map[rune]int, len(s))
		for i := range m {
			m[i] = make(map[rune]int)
		}
		add(s)
	}
	for r.Next() {
		s := r.String()
		if len(s) != len(m) {
			return nil, fmt.Errorf("word length mismatch: started %v, found %v", len(m), len(s))
		}
		add(s)
	}
	if err := r.Err(); err != nil {
		return m, fmt.Errorf("reader error: %v", err)
	}
	return m, nil
}

type ColumnDecoder func (map[rune]int) (rune, int)

func Mode(c map[rune]int) (rune, int) {
	r, m := '?', 0
	for s, n := range c {
		if n > m {
			r, m = s, n
		}
	}
	return r, m
}

func InvMode(c map[rune]int) (rune, int) {
	r, m := '?', int(^uint(0) >> 1)
	for s, n := range c {
		if n < m {
			r, m = s, n
		}
	}
	return r, m
}

type WordDecoder func (r Reader) app.Solution

func NewColumnDecoder(d ColumnDecoder) WordDecoder {
	return func (r Reader) app.Solution {
		m, err := newColumnMap(r)
		if err != nil {
			return app.NewError(err)
		}
		a := make([]rune, len(m))
		for i := range m {
			a[i], _ = d(m[i])
		}
		return app.String(a)
	}
}

func NewSolver(d WordDecoder) app.SolverFunc {
	return func (r io.Reader) app.Solution {
		sc := NewScanner(r)
		return d(sc)
	}
}