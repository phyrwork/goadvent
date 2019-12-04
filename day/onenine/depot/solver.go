package depot

import (
	"fmt"
	"github.com/phyrwork/goadvent/app"
	"io"
	"io/ioutil"
	"strconv"
)

var Rules1 = []Rule{
	NewLengthRule(6),
	NewAdjacentLeastRule(2),
	DecreaseRule,
}

var Rules2 = []Rule{
	NewLengthRule(6),
	NewAdjacentExactRule(2),
	DecreaseRule,
}

func Solve(r io.Reader, rules ...Rule) app.Solution {
	w, err := ioutil.ReadAll(r)
	if err != nil {
		return app.Errorf("read error: %v", err)
	}
	var a, b int
	n, err := fmt.Sscanf(string(w), "%d-%d", &a, &b)
	if err != nil {
		return app.Errorf("scan error: %v", err)
	}
	if n != 2 {
		return app.Errorf("scan arg count error: want %v, got %v", 2, n)
	}
	c := NewRange(a, b, 50)
	s := make(chan string, 50)
	go func () {
		for d := range c {
			s <- strconv.Itoa(d)
		}
		close(s)
	}()
	ok := Filter(s, 50, rules...)
	n = 0
	for range ok {
		n++
	}
	return app.Int(n)
}

func NewSolver(rules ...Rule) app.SolverFunc {
	return func (r io.Reader) app.Solution {
		return Solve(r, rules...)
	}
}
