package abba

import (
	"bufio"
	"github.com/phyrwork/goadvent/app"
	"io"
	"strconv"
)

const (
	HyperStart = '['
	HyperEnd   = ']'
)

type Address string

func (a Address) SupportsTLS() bool {
	ab := false
	it := NewSequencer(a, 4)
	for it.Next() {
		if !it.Sequence().ABBA() {
			continue
		}
		if it.Hyper() {
			return false
		}
		ab = true
	}
	return ab
}

func (a Address) SupportsSSL() bool {
	aba := make(map[string]struct{})
	bab := make(map[string]struct{})
	it := NewSequencer(a, 3)
	for it.Next() {
		seq := it.Sequence()
		if !seq.ABA() {
			continue
		}
		switch {
		case it.Super():
			aba[string(seq)] = struct{}{}
		case it.Hyper():
			bab[string(seq)] = struct{}{}
		}
	}
	for k := range aba {
		a := Sequence(k)
		b := a.BAB()
		if _, ok := bab[string(b)]; ok {
			return true
		}
	}
	return false
}

// TODO: make variable length
type Sequence []rune

func (s Sequence) ABBA() bool {
	if len(s) != 4 {
		return false
	}
	return s[0] == s[3] && s[1] == s[2] && s[0] != s[1]
}

func (s Sequence) ABA() bool {
	if len(s) != 3 {
		return false
	}
	return s[0] == s[2] && s[0] != s[1]
}

func (s Sequence) BAB() Sequence {
	if !s.ABA() {
		return nil
	}
	return Sequence{s[1], s[0], s[1]}
}

type Sequencer struct {
	addr  Address
	hyper bool
	cur   int
	len   int
	seq   Sequence
}

func NewSequencer(addr Address, len int) *Sequencer {
	if len <= 0 {
		return nil
	}
	return &Sequencer{
		addr: addr,
		hyper: false,
		cur:  0,
		len: len,
		seq:  Sequence{},
	}
}

func (s *Sequencer) Next() bool {
	for s.cur + s.len <= len(s.addr) {
		// extract sequence
		w := []rune(string(s.addr))[s.cur:s.cur+s.len]
		s.cur++
		// toggle is in hypernet
		switch w[0] {
		case HyperStart:
			s.hyper = true
		case HyperEnd:
			s.hyper = false
		}
		// can't span hypernet boundary
		span := func (w []rune) bool {
			for _, c := range w {
				switch c {
				case '[', ']':
					return true
				}
			}
			return false
		}
		if span(w) {
			continue
		}
		// yield
		s.seq = w
		return true
	}
	return false
}

func (s *Sequencer) Sequence() Sequence { return s.seq }

func (s *Sequencer) Hyper() bool { return s.hyper }

func (s *Sequencer) Super() bool { return !s.hyper }

type AddressFilter func (Address) bool

func Solve(r io.Reader, f AddressFilter) (string, error) {
	n := 0
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		addr := Address(sc.Text())
		if f(addr) {
			n++
		}
	}
	if err := sc.Err(); err != nil {
		return "", err
	}
	return strconv.Itoa(n), nil
}

func NewSolver(f AddressFilter) app.SolverFunc {
	return func (r io.Reader) (string, error) {
		return Solve(r, f)
	}
}