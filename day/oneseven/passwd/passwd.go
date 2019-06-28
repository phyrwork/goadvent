package passwd

import (
	"bufio"
	"github.com/phyrwork/goadvent/app"
	"github.com/phyrwork/goadvent/iterator"
	"io"
	"strconv"
	"strings"
)

type Passphrase string

func(p Passphrase) Words() iterator.Iterator {
	it := iterator.NewScannerIterator(strings.NewReader(string(p)))
	it.Split(bufio.ScanWords)
	return it
}

func UniqWords(p Passphrase) (bool, error) {
	m := make(map[string]struct{})
	it := p.Words()
	for it.Next() {
		w := it.Value().(string)
		if _, ok := m[w]; ok {
			return false, nil
		}
		m[w] = struct{}{}
	}
	if err := it.Err(); err != nil {
		return false, err
	}
	return true, nil
}

func UniqAnagrams(p Passphrase) (bool, error) {
	m := make(map[int]map[rune]int)
	it := p.Words()
	if err := iterator.Each(it, func (v interface{}) error {
		s := v.(string)
		c := make(map[rune]int)
		for _, r := range s {
			n := c[r] // 0 if not exists
			c[r] = n + 1
		}
		m[len(m)] = c
		return nil
	}); err != nil {
		return false, err
	}
	for s, a := range m {
		for t, b := range m {
			if s == t { // don't compare with self
				continue
			}
			eq := func (a, b map[rune]int) bool {
				if len(a) != len(b) {
					return false
				}
				for r, n := range a {
					m, ok := b[r]
					if !ok || m != n {
						return false
					}
				}
				return true
			}
			if eq(a,b) {
				return false, nil
			}
		}
	}
	return true, nil
}

type ValidFunc func (Passphrase) (bool, error)

func NewSolver(f ValidFunc) app.SolverFunc {
	return func (r io.Reader) (string, error) {
		it := iterator.NewScannerIterator(r)
		it.Split(bufio.ScanLines)
		n := 0
		if err := iterator.Each(it, func (v interface{}) error {
			p := Passphrase(v.(string))
			ok, err := f(p)
			if err != nil {
				return err
			}
			if ok {
				n++
			}
			return nil
		}); err != nil {
			return "", err
		}
		return strconv.Itoa(n), nil
	}
}
