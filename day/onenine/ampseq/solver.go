package ampseq

import (
	"fmt"
	"github.com/phyrwork/goadvent/app"
	"github.com/phyrwork/goadvent/day/onenine/intcode"
	"io"
)

func NewLinear(b int, p intcode.Program, in int, ph ...int) (Sequence, <-chan int) {
	s := NewSequence(p, ph...)
	if len(s) == 0 {
		return nil, nil
	}
	a := s[0]
	a.Read(func () int {
		return in
	})
	w := make(chan int, b)
	a.Write(func (v int) {
		w <- v
	})
	o := w
	for i := 1; i < len(s); i++ {
		r, w := o, make(chan int, b)
		a := s[i]
		a.Read(func () int {
			return <-r
		})
		a.Write(func (v int) {
			w <- v
		})
		o = w
	}
	return s, o
}

func NewFeedback(b int, p intcode.Program, in int, ph ...int) (Sequence, <-chan int) {
	s := NewSequence(p, ph...)
	if len(s) == 0 {
		return nil, nil
	}
	a := s[0]
	w := make(chan int, b)
	a.Write(func (v int) {
		w <- v
	})
	o := w
	for i := 1; i < len(s); i++ {
		r, w := o, make(chan int, b)
		a := s[i]
		a.Read(func () int {
			return <-r
		})
		a.Write(func (v int) {
			w <- v
		})
		o = w
	}
	a.Read(func () int {
		return <-o
	})
	go func () {
		o <-in
	}()
	return s, o
}

type SequenceFunc func (int, intcode.Program, int, ...int) (Sequence, <-chan int)

func Solve(g SequenceFunc, p intcode.Program, ph ...int) (int, []int, error) {
	max := struct {
		sig int
		ph []int
	}{0, nil}
	if err := Perm(ph, func (ph []int) error {
		seq, out := g(20, p, 0, ph...)
		if err := seq.Run(); err != nil {
			return fmt.Errorf("sequence run error: %v", err)
		}
		var sig int
		select {
		case sig = <-out:
		default:
			return fmt.Errorf("sequence has no output")
		}
		if sig > max.sig {
			max.sig = sig
			max.ph = make([]int, len(ph))
			copy(max.ph, ph)
		}
		return nil
	}); err != nil {
		return 0, nil, fmt.Errorf("perm error: %v", err)
	}
	return max.sig, max.ph, nil
}

func NewSolver(g SequenceFunc, ph ...int) app.SolverFunc {
	return func (r io.Reader) app.Solution {
		p, err := intcode.Read(r)
		if err != nil {
			return app.Errorf("program read error: %v", err)
		}
		sig, _, err := Solve(g, p, ph...)
		if err != nil {
			return app.Errorf("solve error: %v", err)
		}
		return app.Int(sig)
	}
}
