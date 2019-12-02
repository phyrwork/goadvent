package checksum

import (
	"bytes"
	"fmt"
	"github.com/phyrwork/goadvent/app"
	"github.com/phyrwork/goadvent/iterator"
	"io"
	"io/ioutil"
)

const (
	MaxInt = int(^uint(0) >> 1)
)

func Max(it iterator.Iterator) (int, error) {
	// TODO: by inspection it is >= 0 although there's nothing in the specification that says this
	r := 0
	for it.Next() {
		if s := it.Value().(int); s > r {
			r = s
		}
	}
	return r, it.Err()
}

func Min(it iterator.Iterator) (int, error) {
	// TODO: by inspection it is >= 0 although there's nothing in the specification that says this
	r := MaxInt
	for it.Next() {
		if s := it.Value().(int); s < r {
			r = s
		}
	}
	return r, it.Err()
}

func Diff(it iterator.Iterator) (int, error) {
	v, err := iterator.Array(it)
	if err != nil {
		return 0, err
	}
	it = iterator.NewArrayIterator(v...)
	a, err := Min(it)
	if err != nil {
		return a, err
	}
	if err := it.(iterator.ResetIterator).Reset(); err != nil {
		return 0, err
	}
	b, err := Max(it)
	if err != nil {
		return b, err
	}
	return b - a, nil
}

func FactorDiv(it iterator.Iterator) (int, error) {
	m := make(map[int]struct{})
	for it.Next() {
		m[it.Value().(int)] = struct{}{}
	}
	if err := it.Err(); err != nil {
		return 0, nil
	}
	var a, b int
	n := 0
	for c := range m {
		for d := range m {
			if c == d {
				continue
			}
			if c % d == 0 {
				a, b = c, d
				n++
			}
		}
	}
	switch n {
	case 0:
		return 0, fmt.Errorf("no divisible pairs")
	case 1:
		return a/b, nil
	default:
		return 0, fmt.Errorf("multiple (%v) divisible pairs", n)
	}
}

type SumFunc func(iterator.Iterator) (int, error)

func Checksum(it iterator.Iterator, sum SumFunc) (int, error) {
	s := 0
	for it.Next() {
		it := it.Value().(iterator.Iterator)
		d, err := sum(it)
		if err != nil {
			return s, err
		}
		s += d
	}
	return s, it.Err()
}

func NewSolver(sum SumFunc) app.Solver {
	return app.SolverFunc(func (r io.Reader) app.Solution{
		b, err := ioutil.ReadAll(r)
		if err != nil {
			return app.NewError(err)
		}
		r = bytes.NewReader(b)
		sc := NewRowScanner(r)
		c, err := Checksum(sc, sum)
		if err != nil {
			return app.NewError(err)
		}
		return app.Int(c)
	})
}

