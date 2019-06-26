package captcha

import (
	"bufio"
	"container/ring"
	"fmt"
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

func (d *Digits) Sum() int {
	s := 0
	d.each(func (r *ring.Ring) {
		a, b := r.Value.(int), r.Next().Value.(int)
		if a == b {
			s += a
		}
	})
	return s
}