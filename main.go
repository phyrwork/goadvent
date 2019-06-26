package main

import (
	"fmt"
	"github.com/phyrwork/goadvent/day/oneseven/captcha"
	"io"
	"log"
	"os"
	"strconv"
)

type Solver interface {
	Solve(rd io.Reader) (string, error)
}

type SolverFunc func (io.Reader) (string, error)

func (f SolverFunc) Solve(rd io.Reader) (string, error) { return f(rd) }

var solvers = map[string]Solver{
	"2017.1.1": SolverFunc(func (rd io.Reader) (string, error) {
		d, err := captcha.NewDigits(rd)
		if err != nil {
			return "", err
		}
		i := d.Sum()
		return strconv.Itoa(i), nil
	}),
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatalln("error: solver name not given")
	}
	id := args[0]
	solver := solvers[id]
	if solver == nil {
		log.Fatalf("error: solver '%v' not found", id)
	}
	out, err := solver.Solve(os.Stdin)
	if err != nil {
		log.Fatalf("solver error: %v", err)
	}
	fmt.Println(out)
}