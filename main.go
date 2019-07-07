package main

import (
	"fmt"
	"github.com/phyrwork/goadvent/app"
	"github.com/phyrwork/goadvent/day/oneseven/captcha"
	"github.com/phyrwork/goadvent/day/oneseven/checksum"
	"github.com/phyrwork/goadvent/day/oneseven/hex"
	"github.com/phyrwork/goadvent/day/oneseven/jump"
	"github.com/phyrwork/goadvent/day/oneseven/knot"
	"github.com/phyrwork/goadvent/day/oneseven/passwd"
	"github.com/phyrwork/goadvent/day/oneseven/stream"
	"github.com/phyrwork/goadvent/day/onesix/taxi"
	"log"
	"os"
)

var solvers = map[string]app.Solver{
	"2016.1.1": app.SolverFunc(taxi.SolveDist),
	"2017.1.1": captcha.NewSolver(captcha.Next),
	"2017.1.2": captcha.NewSolver(captcha.Half),
	"2017.2.1": checksum.NewSolver(checksum.Diff),
	"2017.2.2": checksum.NewSolver(checksum.FactorDiv),
	"2017.4.1": passwd.NewSolver(passwd.UniqWords),
	"2017.4.2": passwd.NewSolver(passwd.UniqAnagrams),
	"2017.5.1": jump.NewSolver(jump.Jump),
	"2017.5.2": jump.NewSolver(jump.StrangeJump),
	"2017.9.1": stream.NewSolver(func (g *stream.Group) int { return g.Score(1) }),
	"2017.9.2": stream.NewSolver(func (g *stream.Group) int { return g.Chars() }),
	"2017.10.1": app.SolverFunc(knot.SolveSparse),
	"2017.10.2": app.SolverFunc(knot.KnotHash),
	"2017.11.1": hex.NewSolver(hex.Sum),
	"2017.11.2": hex.NewSolver(hex.MaxDist),
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