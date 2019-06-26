package app

import "io"

type Solver interface {
	Solve(rd io.Reader) (string, error)
}

type SolverFunc func (io.Reader) (string, error)

func (f SolverFunc) Solve(rd io.Reader) (string, error) { return f(rd) }