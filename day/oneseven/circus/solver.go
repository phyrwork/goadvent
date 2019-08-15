package circus

import (
	"fmt"
	"github.com/phyrwork/goadvent/app"
	"io"
	"strconv"
)

func NewSolver(f func (Circus) (string, error)) app.SolverFunc {
	return func(r io.Reader) (string, error) {
		descs, err := Parse(r)
		if err != nil {
			return "", err
		}
		c, err := NewCircus(descs...)
		if err != nil {
			return "", err
		}
		return f(c)
	}
}

func SolveBase(c Circus) (string, error) {
	bt, err := c.Base()
	if err != nil {
		return "", err
	}
	return bt.Name, nil
}

func SolveWeight(c Circus) (string, error) {
	mod, err := c.Balance()
	if err != nil {
		return "", err
	}
	if len(mod) != 1 {
		return "", fmt.Errorf("more than one tower modified")
	}
	var w int
	for _, w = range mod {
		break
	}
	return strconv.Itoa(w), nil
}
