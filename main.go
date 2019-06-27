package main

import (
	"fmt"
	"github.com/phyrwork/goadvent/app"
	"github.com/phyrwork/goadvent/day/oneseven/captcha"
	"github.com/phyrwork/goadvent/day/oneseven/checksum"
	"log"
	"os"
)



var solvers = map[string]app.Solver{
	"2017.1.1": captcha.NewSolver(captcha.Next),
	"2017.1.2": captcha.NewSolver(captcha.Half),
	"2017.2.1": checksum.NewSolver(checksum.Diff),
	"2017.2.2": checksum.NewSolver(checksum.FactorDiv),
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