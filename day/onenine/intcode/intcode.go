package intcode

import (
	"bufio"
	"fmt"
	"github.com/phyrwork/goadvent/app"
	"github.com/phyrwork/goadvent/iterator"
	"io"
	"strconv"
)

type Program []int

type Memory Program

type Op func (m Memory, pci int) (pco int, ok bool)

type Machine struct {
	m   Memory
	pc  int
	op  map[int]Op
	err error
}

func (m *Machine) Next() bool {
	if m.pc >= len(m.m) {
		m.err = fmt.Errorf("pc out of bounds: %v/%v", m.pc, len(m.m))
		return false
	}
	oc := m.m[m.pc]
	op, ok := m.op[oc]
	if !ok {
		m.err = fmt.Errorf("unknown opcode: %v", oc)
		return false
	}
	m.pc, ok = op(m.m, m.pc)
	return ok
}

func (m *Machine) Err() error { return m.err }

var DefaultOps = map[int]Op {
	1: func (m Memory, pc int) (int, bool) {
		a, b, o := m[pc+1], m[pc+2], m[pc+3]
		m[o] = m[a] + m[b]
		return pc + 4, true
	},
	2: func (m Memory, pc int) (int, bool) {
		a, b, o := m[pc+1], m[pc+2], m[pc+3]
		m[o] = m[a] * m[b]
		return pc + 4, true
	},
	99: func (m Memory, pc int) (int, bool) {
		return pc + 1, false
	},
}

func Read(r io.Reader) (Program, error) {
	p := make(Program, 0)
	sc := bufio.NewScanner(r)
	sc.Split(iterator.SplitComma)
	for sc.Scan() {
		s := sc.Text()
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, fmt.Errorf("atoi error: %v", err)
		}
		p = append(p, i)
	}
	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("scan error: %v", err)
	}
	return p, nil
}

func Pair(_p Program, n, v int) (int, error) {
	p := make(Program, len(_p))
	copy(p, _p)
	p[1] = n
	p[2] = v
	m := Machine{
		m:  Memory(p),
		op: DefaultOps,
	}
	for m.Next() {

	}
	if err := m.Err(); err != nil {
		return 0, fmt.Errorf("machine error: %v", err)
	}
	return m.m[0], nil
}

func Solve1(r io.Reader) app.Solution {
	p, err := Read(r)
	if err != nil {
		return app.Errorf("read error: %v", err)
	}
	o, err := Pair(p, 12,  2)
	if err != nil {
		return app.Errorf("pair error: %v", err)
	}
	return app.Int(o)
}

type pair struct {
	n, v int
}

func Solve2(r io.Reader) app.Solution {
	p, err := Read(r)
	if err != nil {
		return app.Errorf("read error: %v", err)
	}
	c := make(chan pair)
	go func () {
		for n := 0; n <= 99; n++ {
			for v := 0; v <= 99; v++ {
				c <-pair{n, v}
			}
		}
		close(c)
	}()
	for nv := range c {
		o, err := Pair(p, nv.n, nv.v)
		if err != nil {
			continue
		}
		if o == 19690720 {
			return app.Int(100 * nv.n + nv.v)
		}
	}
	return app.Errorf("solution not found")
}