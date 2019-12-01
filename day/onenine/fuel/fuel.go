package fuel

import (
	"bufio"
	"fmt"
	"github.com/phyrwork/goadvent/app"
	"io"
	"strconv"
)

type FuelFunc func (int) int

type Module int

func NewModule(w int) Module { return Module(w) }

func (m Module) Weight() int { return int(m) }

func Read(r io.Reader) ([]Module, error) {
	m := make([]Module, 0)
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		s := sc.Text()
		w, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("module weight decode error: %v", err)
		}
		m = append(m, NewModule(w))
	}
	return m, nil
}

func Fuel(m []Module, f FuelFunc) int {
	s := 0
	for _, m := range m {
		s += f(m.Weight())
	}
	return s
}

func ModuleFuel(w int) int { return w/3 - 2 }

func RocketFuel(w int) int {
	s := 0
	f := ModuleFuel(w)
	s += f
	for f > 0 {
		f = ModuleFuel(f)
		if f < 0 {
			f = 0
		}
		s += f
	}
	return s
}

func NewSolver(f FuelFunc) app.SolverFunc {
	return func (r io.Reader) (string, error) {
		m, err := Read(r)
		if err != nil {
			return "", fmt.Errorf("parse error")
		}
		f := Fuel(m, f)
		return strconv.Itoa(f), nil
	}
}
