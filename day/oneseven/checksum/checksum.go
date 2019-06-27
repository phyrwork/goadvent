package checksum

import (
	"bytes"
	"fmt"
	"github.com/phyrwork/goadvent/app"
	"io"
	"io/ioutil"
	"strconv"
)

const (
	MaxInt = int(^uint(0) >> 1)
)

type Values interface {
	Next() bool
	Int() int
	Err() error
	Reset() error
}

type Rows interface {
	Next() bool
	Row() Values
	Err() error
	Reset() error
}

func Max(v Values) (int, error) {
	// TODO: by inspection r is >= 0 although there's nothing in the specification that says this
	r := 0
	if err := v.Reset(); err != nil {
		return r, err
	}
	for v.Next() {
		if s := v.Int(); s > r {
			r = s
		}
	}
	return r, v.Err()
}

func Min(v Values) (int, error) {
	// TODO: by inspection r is < max(int) although there's nothing in the specification that says this
	r := MaxInt
	if err := v.Reset(); err != nil {
		return r, err
	}
	for v.Next() {
		if s := v.Int(); s < r {
			r = s
		}
	}
	return r, v.Err()
}

func Diff(v Values) (int, error) {
	a, err := Min(v)
	if err != nil {
		return a, err
	}
	b, err := Max(v)
	if err != nil {
		return b, err
	}
	return b - a, nil
}

func FactorDiv(v Values) (int, error) {
	if err := v.Reset(); err != nil {
		return 0, err
	}
	m := make(map[int]struct{})
	for v.Next() {
		m[v.Int()] = struct{}{}
	}
	if err := v.Err(); err != nil {
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

type SumFunc func(Values) (int, error)

func Checksum(r Rows, sum SumFunc) (int, error) {
	s := 0
	if err := r.Reset(); err != nil {
		return 0, err
	}
	for r.Next() {
		v := r.Row()
		d, err := sum(v)
		if err != nil {
			return s, err
		}
		s += d
	}
	return s, r.Err()
}

func NewSolver(sum SumFunc) app.Solver {
	return app.SolverFunc(func (r io.Reader) (string, error){
		b, err := ioutil.ReadAll(r)
		if err != nil {
			return "", err
		}
		r = bytes.NewReader(b)
		sc := NewRowScanner(r)
		c, err := Checksum(sc, sum)
		if err != nil {
			return "", err
		}
		return strconv.Itoa(c), nil
	})
}

