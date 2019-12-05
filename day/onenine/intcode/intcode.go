package intcode

import (
	"fmt"
	"strconv"
)

type Program []int

type Memory Program

// TODO: don't need to pass PC anymore
type Op func (m *Machine, pci int) (pco int, ok bool)

const (
	Position = '0'
	Immediate = '1'
)

type ArgModes []rune

func NewArgModes(ov int) ArgModes {
	if ov < 0 {
		panic("negative op val")
	}
	s := strconv.Itoa(ov)
	c := []rune(s)
	if len(c) < 3 {
		return nil
	}
	c = c[:len(c) - 2] // keep only arg chars
	return c
}

func (am ArgModes) Mode(arg int) rune {
	if arg >= len(am) {
		return Position
	}
	// codes are ascending right to left
	return am[len(am) - (1 + arg)]
}

type Machine struct {
	m   Memory
	pc  int
	op  map[int]Op
	am  ArgModes
	err error
	in  func () int
	out func (int)
}

func parseOp(ov int) int {
	return ov % 100
}

func (m *Machine) getArg(pc int, o int) int {
	am := m.am.Mode(o)
	switch am {
	case Immediate:
		return m.m[pc+1+o]
	case Position:
		a := m.m[pc+1+o]
		return m.m[a]
	default:
		panic(fmt.Sprintf("unknown arg mode %v", am))
	}
}

func (m *Machine) Next() bool {
	if m.pc >= len(m.m) {
		m.err = fmt.Errorf("pc out of bounds: %v/%v", m.pc, len(m.m))
		return false
	}
	ov := m.m[m.pc]
	oc := parseOp(ov)
	m.am = NewArgModes(ov)
	op, ok := m.op[oc]
	if !ok {
		m.err = fmt.Errorf("unknown opcode: %v", oc)
		return false
	}
	m.pc, ok = op(m, m.pc)
	return ok
}

func (m *Machine) Err() error { return m.err }

var DefaultOps = map[int]Op {
	// add
	1: func (m *Machine, pc int) (int, bool) {
		a, b, o := m.getArg(pc, 0), m.getArg(pc, 1), m.m[pc+3]
		m.m[o] = a + b
		return pc + 4, true
	},
	// mul
	2: func (m *Machine, pc int) (int, bool) {
		a, b, o := m.getArg(pc, 0), m.getArg(pc, 1), m.m[pc+3]
		m.m[o] = a * b
		return pc + 4, true
	},
	// read
	3: func (m *Machine, pc int) (int, bool) {
		if m.in == nil {
			panic("nil input handler")
		}
		o := m.m[pc+1]
		m.m[o] = m.in()
		return pc + 2, true
	},
	// write
	4: func (m *Machine, pc int) (int, bool) {
		if m.out == nil {
			panic("nil output handler")
		}
		i := m.getArg(pc, 0)
		m.out(i)
		return pc + 2, true
	},
	// halt
	99: func (m *Machine, pc int) (int, bool) {
		return pc + 1, false
	},
}