package spinlock

import (
	"bufio"
	"github.com/phyrwork/goadvent/app"
	"io"
	"log"
	"strconv"
)

func NewSolver(c func () <-chan int, v func (*Spinlock) int) app.SolverFunc {
	return func (r io.Reader) (string, error) {
		sc := bufio.NewScanner(r)
		sc.Split(bufio.ScanWords)
		var s string
		for sc.Scan() {
			s = sc.Text()
			break
		}
		if err := sc.Err(); err != nil {
			return "", err
		}
		step, err := strconv.Atoi(s)
		if err != nil {
			return "", err
		}
		l := NewSpinlock(step)
		l.Stream(c())
		return strconv.Itoa(v(l)), nil
	}
}

func Next(l *Spinlock) int { return l.Peek(1) }

func ZeroNext(l *Spinlock) int {
	o, m := l.Find(0)
	if o < 0 {
		log.Panic("zero not found") // This should never happen
	}
	return Next(m)
}