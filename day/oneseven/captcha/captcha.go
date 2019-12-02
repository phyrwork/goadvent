package captcha

import (
	"bufio"
	"container/ring"
	"fmt"
	"github.com/phyrwork/goadvent/app"
	"io"
	"strconv"
)

type Digits struct {
	r *ring.Ring
}

func NewDigits(rd io.Reader) (*Digits, error) {
	var r *ring.Ring
	sc := bufio.NewScanner(rd)
	sc.Split(bufio.ScanRunes)
	for sc.Scan() {
		s := sc.Text()
		if s[0] == '\n' {
			break
		}
		d, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("error decoding digit %v", err)
		}
		n := ring.New(1)
		n.Value = d
		if r != nil {
			r.Link(n)
		}
		r = n
	}
	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("error scanning digits: %v", err)
	}
	return &Digits{r.Next()}, nil
}

func (d *Digits) String() string {
	s := ""
	d.r.Do(func (v interface{}) {
		i := v.(int)
		c := fmt.Sprintf("%d", i)
		s += c
	})
	return s
}

func (d *Digits) each(f func (*ring.Ring)) {
	r := d.r
	if r == nil {
		return
	}
	c := r
	for {
		f(c)
		c = c.Next()
		if c == r {
			break
		}
	}
}

type CmpFunc func (curr *ring.Ring) *ring.Ring

func Next(r *ring.Ring) *ring.Ring { return r.Next() }

func Half(r *ring.Ring) *ring.Ring {
	d := r.Len()/2
	for i := 0; i < d; i++ {
		r = r.Next()
	}
	return r
}

func (d *Digits) Sum(cmp CmpFunc) int {
	s := 0
	d.each(func (r *ring.Ring) {
		a, b := r.Value.(int), cmp(r).Value.(int)
		if a == b {
			s += a
		}
	})
	return s
}

func NewSolver(cmp CmpFunc) app.SolverFunc {
	return func (rd io.Reader) app.Solution {
		d, err := NewDigits(rd)
		if err != nil {
			return app.NewError(err)
		}
		i := d.Sum(cmp)
		return app.Int(i)
	}
}