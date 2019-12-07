package ampseq

import (
	"fmt"
	"github.com/phyrwork/goadvent/app"
	"github.com/phyrwork/goadvent/day/onenine/intcode"
	"io"
)

func solve1(p intcode.Program) (int, []int, error) {
	ph := []int{0, 1, 2, 3, 4}
	max := struct {
		sig int
		ph []int
	}{0, nil}
	if err := Perm(ph, func (ph []int) error {
		seq := NewSequence(p, 0, ph...)
		sig, err := seq.Run()
		if err != nil {
			return fmt.Errorf("sequence run error: %v", err)
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

func Solve1(r io.Reader) app.Solution {
	p, err := intcode.Read(r)
	if err != nil {
		return app.Errorf("program read error: %v", err)
	}
	sig, _, err := solve1(p)
	if err != nil {
		return app.Errorf("solve error: %v", err)
	}
	return app.Int(sig)
}
