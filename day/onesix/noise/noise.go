package noise

import (
	"fmt"
	"io"
)

type DecodeFunc func (r Reader) (string, error)

func DecodeMode(r Reader) (string, error) {
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
			return "", fmt.Errorf("word length mismatch: started %v, found %v", len(m), len(s))
		}
		add(s)
	}
	if err := r.Err(); err != nil {
		return "", fmt.Errorf("reader error: %v", err)
	}
	max := func (c map[rune]int) (rune, int) {
		r, m := '?', 0
		for s, n := range c {
			if n > m {
				r, m = s, n
			}
		}
		return r, m
	}
	a := make([]rune, len(m))
	for i := range m {
		a[i], _ = max(m[i])
	}
	return string(a), nil
}

func Solve(r io.Reader) (string, error) {
	sc := NewScanner(r)
	return DecodeMode(sc)
}