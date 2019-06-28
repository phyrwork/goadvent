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

type JumpFunc func (pc, op *int)

type Jumper struct {
	p Program
	pc int
	jmp JumpFunc
}

func NewJumper(p Program, jmp JumpFunc) *Jumper { return &Jumper{p, 0, jmp} }

func (j *Jumper) OK() bool { return j.pc < len(j.p) }

func (j *Jumper) Exec() {
	op := &j.p[j.pc]
	j.jmp(&j.pc, op)
}

func Jump(pc, op *int) {
	d := *op
	*op++
	*pc += d
}

func StrangeJump(pc, op *int) {
	d := *op
	if d >= 3 {
		*op--
	} else {
		*op++
	}
	*pc += d
}

func Exec(exe Executable) int {
	c := 0
	for exe.OK() {
		exe.Exec()
		c++
	}
	return c
}

func NewSolver(jmp JumpFunc) app.SolverFunc {
	return app.SolverFunc(func (r io.Reader) (string, error) {
		p, err := ReadProgram(r)
		if err != nil {
			return "", err
		}
		exe := NewJumper(p, jmp)
		c := Exec(exe)
		return strconv.Itoa(c), nil
	})
}