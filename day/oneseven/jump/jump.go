package jump

import (
	"bufio"
	"github.com/phyrwork/goadvent/app"
	"github.com/phyrwork/goadvent/iterator"
	"io"
	"strconv"
)

type Program []int

func ReadProgram(r io.Reader) (Program, error) {
	it := iterator.NewScannerIterator(r)
	it.Split(bufio.ScanLines)
	p := make(Program, 0)
	return p, iterator.Each(it, func (v interface{}) error {
		s := v.(string)
		d, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		p = append(p, d)
		return nil
	})
}

type Executable interface {
	OK() bool
	Exec()
}

type Jumper struct {
	p Program
	pc int
}

func NewJumper(p Program) *Jumper { return &Jumper{p, 0} }

func (j *Jumper) OK() bool { return j.pc < len(j.p) }

func (j *Jumper) Exec() {
	op := &j.p[j.pc]
	d := *op
	*op += 1
	j.pc += d
}

func Exec(exe Executable) int {
	c := 0
	for exe.OK() {
		exe.Exec()
		c++
	}
	return c
}

func NewSolver(f func (Program) Executable) app.SolverFunc {
	return app.SolverFunc(func (r io.Reader) (string, error) {
		p, err := ReadProgram(r)
		if err != nil {
			return "", err
		}
		exe := f(p)
		c := Exec(exe)
		return strconv.Itoa(c), nil
	})
}