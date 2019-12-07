package ampseq

import (
	"fmt"
	"github.com/phyrwork/goadvent/day/onenine/intcode"
)

type Amplifier struct {
	m *intcode.Machine
	ph int
	sig int
}

func NewAmplifier(p intcode.Program, ph int) *Amplifier {
	a := &Amplifier{
		m:   intcode.NewMachine(intcode.DefaultOps, p),
		ph:  ph,
		sig: 0,
	}
	a.m.Write = func (sig int) {
		a.sig = sig
	}
	return a
}

func (a *Amplifier) Read(sig int) {
	in := []int{a.ph, sig}
	a.m.Read = func () int {
		if len(in) == 0 {
			return 0
		}
		v := in[0]
		in = in[1:]
		return v
	}
}

func (a *Amplifier) Run() (int, error) {
	for a.m.Next() {
	}
	if err := a.m.Err(); err != nil {
		return 0, fmt.Errorf("run error: %v", err)
	}
	return a.sig, nil
}

type Sequence []*Amplifier

func NewSequence(p intcode.Program, in int, ph ...int) Sequence {
	if len(ph) == 0 {
		return nil
	}
	a := make([]*Amplifier, len(ph))
	for i, ph := range ph {
		a[i] = NewAmplifier(p, ph)
	}
	return a
}

func (s Sequence) Run() (int, error) {
	sig := 0
	for _, a := range s {
		a.Read(sig)
		var err error
		sig, err = a.Run()
		if err != nil {
			return 0, err
		}
	}
	return sig, nil
}

// Perm calls f with each permutation of a.
func Perm(a []int, f func([]int) error) error {
	return perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []int, f func([]int) error, i int) error {
	if i > len(a) {
		return f(a)
	}
	if err := perm(a, f, i+1); err != nil {
		return err
	}
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		if err := perm(a, f, i+1); err != nil {
			return err
		}
		a[i], a[j] = a[j], a[i]
	}
	return nil
}