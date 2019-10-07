package main

import (
	"fmt"
	"github.com/phyrwork/goadvent/app"
	"github.com/phyrwork/goadvent/day/onefive/houses"
	"github.com/phyrwork/goadvent/day/oneseven/captcha"
	"github.com/phyrwork/goadvent/day/oneseven/checksum"
	"github.com/phyrwork/goadvent/day/oneseven/circus"
	"github.com/phyrwork/goadvent/day/oneseven/hex"
	"github.com/phyrwork/goadvent/day/oneseven/jump"
	"github.com/phyrwork/goadvent/day/oneseven/knot"
	"github.com/phyrwork/goadvent/day/oneseven/passwd"
	"github.com/phyrwork/goadvent/day/oneseven/plumber"
	"github.com/phyrwork/goadvent/day/oneseven/realloc"
	"github.com/phyrwork/goadvent/day/oneseven/registers"
	"github.com/phyrwork/goadvent/day/oneseven/spinlock"
	"github.com/phyrwork/goadvent/day/oneseven/stream"
	"github.com/phyrwork/goadvent/day/onesix/abba"
	"github.com/phyrwork/goadvent/day/onesix/chess"
	"github.com/phyrwork/goadvent/day/onesix/explosive"
	"github.com/phyrwork/goadvent/day/onesix/keypad"
	"github.com/phyrwork/goadvent/day/onesix/kiosk"
	"github.com/phyrwork/goadvent/day/onesix/noise"
	"github.com/phyrwork/goadvent/day/onesix/taxi"
	"log"
	"os"
)

var solvers = map[string]app.Solver{
	"2015.3.1": houses.NewSolver(houses.CountUnique),
	"2016.1.1": taxi.NewSolver(taxi.Walk),
	"2016.1.2": taxi.NewSolver(taxi.WalkUntilRevisit),
	"2016.2.1": keypad.NewSolver(keypad.SquareKeypad, keypad.Position{1, 1}),
	"2016.2.2": keypad.NewSolver(keypad.DiamondKeypad, keypad.Position{0, 2}),
	"2016.4.1": app.SolverFunc(kiosk.SolveSumReal),
	"2016.4.2": app.SolverFunc(kiosk.SolveNorthPoleRoom),
	"2016.5.1": app.SolverFunc(chess.SolveAppend),
	"2016.5.2": app.SolverFunc(chess.SolveFiller),
	"2016.6.1": noise.NewSolver(noise.NewColumnDecoder(noise.Mode)),
	"2016.6.2": noise.NewSolver(noise.NewColumnDecoder(noise.InvMode)),
	"2016.7.1": abba.NewSolver(abba.Address.SupportsTLS),
	"2016.7.2": abba.NewSolver(abba.Address.SupportsSSL),
	"2016.9.1": explosive.NewSolver(false),
	"2016.9.2": explosive.NewSolver(true),
	"2017.1.1": captcha.NewSolver(captcha.Next),
	"2017.1.2": captcha.NewSolver(captcha.Half),
	"2017.2.1": checksum.NewSolver(checksum.Diff),
	"2017.2.2": checksum.NewSolver(checksum.FactorDiv),
	"2017.4.1": passwd.NewSolver(passwd.UniqWords),
	"2017.4.2": passwd.NewSolver(passwd.UniqAnagrams),
	"2017.5.1": jump.NewSolver(jump.Jump),
	"2017.5.2": jump.NewSolver(jump.StrangeJump),
	"2017.6.1": app.SolverFunc(realloc.SolveFirst),
	"2017.6.2": app.SolverFunc(realloc.SolveSize),
	"2017.7.1": circus.NewSolver(circus.SolveBase),
	"2017.7.2": circus.NewSolver(circus.SolveWeight),
	"2017.8.1": app.SolverFunc(registers.SolveMaxEnd),
	"2017.8.2": app.SolverFunc(registers.SolveMaxAny),
	"2017.9.1": stream.NewSolver(func (g *stream.Group) int { return g.Score(1) }),
	"2017.9.2": stream.NewSolver(func (g *stream.Group) int { return g.Chars() }),
	"2017.10.1": app.SolverFunc(knot.SolveSparse),
	"2017.10.2": app.SolverFunc(knot.KnotHash),
	"2017.11.1": hex.NewSolver(hex.Sum),
	"2017.11.2": hex.NewSolver(hex.MaxDist),
	"2017.12.1": app.SolverFunc(plumber.SolveSize),
	"2017.12.2": app.SolverFunc(plumber.SolveCount),
	"2017.17.1": spinlock.NewSolver(func () <-chan int { return spinlock.NewSequence(1, 2017) }, spinlock.Next),
	"2017.17.2": spinlock.NewSolver(func () <-chan int { return spinlock.NewSequence(1, 50000000) }, spinlock.ZeroNext),
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