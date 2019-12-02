package circus

import (
	"github.com/phyrwork/goadvent/app"
	"io"
)

func NewSolver(f func (Circus) app.Solution) app.SolverFunc {
	return func(r io.Reader) app.Solution {
		descs, err := Parse(r)
		if err != nil {
			return app.NewError(err)
		}
		c, err := NewCircus(descs...)
		if err != nil {
			return app.NewError(err)
		}
		return f(c)
	}
}

func SolveBase(c Circus) app.Solution {
	bt, err := c.Base()
	if err != nil {
		return app.NewError(err)
	}
	return app.String(bt.Name)
}

func SolveWeight(c Circus) app.Solution {
	mod, err := c.Balance()
	if err != nil {
		return app.NewError(err)
	}
	if len(mod) != 1 {
		return app.Errorf("more than one tower modified")
	}
	var w int
	for _, w = range mod {
		break
	}
	return app.Int(w)
}
