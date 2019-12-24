package intcode2

import (
	"fmt"
	"log"
	"strconv"
)

type Reader interface {
	Read() int
}

type ReadFunc func () int

func (f ReadFunc) Read() int { return f() }

type Writer interface {
	Write(v int)
}

type WriteFunc func (v int)

func (f WriteFunc) Write(v int) {
	f(v)
}

type Memory interface {
	Load(o uint) int
	Store(o uint, v int)
}

type MapMemory map[uint]int

func (m MapMemory) Load(o uint) int {
	return m[o]
}

func (m MapMemory) Store(o uint, v int) {
	m[o] = v
}

type Immediate int

func (p Immediate) Read() int {
	return int(p)
}

type Refer struct {
	Mem    Memory
	Offset uint
}

func (ref Refer) Read() int {
	return ref.Mem.Load(ref.Offset)
}

func (ref Refer) Write(v int) {
	ref.Mem.Store(ref.Offset, v)
}

type Opcode uint

func (c Opcode) ID() uint {
	return uint(c) % 100
}

func (c Opcode) Modes() map[uint]uint {
	s := strconv.Itoa(int(c))
	w := []rune(s)
	m := make(map[uint]uint)
	for i := 2; i < len(w); i++ {
		m[uint(i-2)] = uint(w[len(w)-1-i]) - '0'
	}
	return m
}

type Op func (v ...interface{})

type Proto struct {
	Op   Op
	Narg uint
}

type Instruct struct {
	Op   Op
	Args []interface{}
}

func (inst Instruct) Exec() {
	inst.Op(inst.Args...)
}

type ParamFunc func (p int) interface{}

type Decoder struct {
	Mem   Memory
	Modes map[uint]ParamFunc
	Ops   map[uint]Proto
}

func (d Decoder) Decode(o uint) Instruct {
	c := Opcode(d.Mem.Load(o))
	id := c.ID()
	proto, ok := d.Ops[id]
	if !ok {
		log.Panicf("unknown proto %v", id)
	}
	args := make([]interface{}, proto.Narg)
	modes := c.Modes()
	for i := uint(0); i < proto.Narg; i++ {
		param := d.Mem.Load(o+1+i)
		mode := modes[i]
		wrap, ok := d.Modes[mode]
		if !ok {
			log.Panicf("unknown mode %v", mode)
		}
		args[i] = wrap(param)
	}
	return Instruct{proto.Op, args}
}

type Fetcher struct {
	pc     *uint
	halted *bool
	dc     *Decoder
	inst   Instruct
}

func (f *Fetcher) Fetch() bool {
	if *f.halted {
		return false
	}
	f.inst = f.dc.Decode(*f.pc)
	*f.pc += 1+uint(len(f.inst.Args))
	return true
}

func (f *Fetcher) Instruct() Instruct {
	return f.inst
}

type ArgCountError struct {
	Op   string
	Want uint
	Got  uint
}

func (err ArgCountError) Error() string {
	return fmt.Sprintf("op %v arg count error: want %v, got %v", err.Op, err.Want, err.Got)
}

var Add = Proto{
	Op: func (v ...interface{}) {
		if len(v) != 3 {
			log.Panic(ArgCountError{"ADD", 3, uint(len(v))})
		}
		a := v[0].(Reader).Read()
		b := v[1].(Reader).Read()
		v[2].(Writer).Write(a+b)
	},
	Narg: 3,
}

var Mul = Proto{
	Op: func (v ...interface{}) {
		if len(v) != 3 {
			log.Panic(ArgCountError{"MUL", 3, uint(len(v))})
		}
		a := v[0].(Reader).Read()
		b := v[1].(Reader).Read()
		v[2].(Writer).Write(a*b)
	},
	Narg: 3,
}

var Less = Proto{
	Op: func (v ...interface{}) {
		if len(v) != 3 {
			log.Panic(ArgCountError{"LT", 3, uint(len(v))})
		}
		a := v[0].(Reader).Read()
		b := v[1].(Reader).Read()
		var r int
		if a < b {
			r = 1
		} else {
			r = 0
		}
		v[2].(Writer).Write(r)
	},
	Narg: 3,
}

var Eq = Proto{
	Op: func (v ...interface{}) {
		if len(v) != 3 {
			log.Panic(ArgCountError{"EQ", 3, uint(len(v))})
		}
		a := v[0].(Reader).Read()
		b := v[1].(Reader).Read()
		var r int
		if a == b {
			r = 1
		} else {
			r = 0
		}
		v[2].(Writer).Write(r)
	},
	Narg: 3,
}

type Machine struct {
	pc     uint
	halted bool
	rb     int
	dc     *Decoder
	fetch  *Fetcher
	mem    Memory
	in     Reader
	out    Writer
}

func NewMachine(mem Memory) *Machine {
	m := &Machine{
		mem: mem,
	}
	m.dc = &Decoder{
		Mem:   mem,
		Modes: map[uint]ParamFunc{
			0: m.pos,
			1: func (p int) interface{} {
				return Immediate(p)
			},
			2: m.rel,
		},
		Ops:   map[uint]Proto{
			1: Add,
			2: Mul,
			3: {m.read, 1},
			4: {m.write, 1},
			5: {m.jmpnz, 2},
			6: {m.jmpz, 2},
			7: Less,
			8: Eq,
			9: {m.rboff, 1},
			99: {m.halt, 0},
		},
	}
	m.fetch = &Fetcher{
		pc:     &m.pc,
		halted: &m.halted,
		dc:     m.dc,
		inst:   Instruct{},
	}
	return m
}

func (m *Machine) pos(p int) interface{} {
	return Refer{
		Mem:    m.mem,
		Offset: uint(p),
	}
}

func (m *Machine) rel(p int) interface{} {
	return Refer{
		Mem:    m.mem,
		Offset: uint(m.rb + p),
	}
}

func (m *Machine) jmp(o uint) {
	m.pc = o
}

func (m *Machine) read(v ...interface{}) {
	if len(v) != 1 {
		log.Panic(ArgCountError{"R", 1, uint(len(v))})
	}
	d := m.in.Read()
	v[0].(Writer).Write(d)
}

func (m *Machine) write(v ...interface{}) {
	if len(v) != 1 {
		log.Panic(ArgCountError{"W", 1, uint(len(v))})
	}
	d := v[0].(Reader).Read()
	m.out.Write(d)
}

func (m *Machine) jmpz(v ...interface{}) {
	if len(v) != 2 {
		log.Panic(ArgCountError{"JMPZ", 2, uint(len(v))})
	}
	a := v[0].(Reader).Read()
	if a == 0 {
		o := v[1].(Reader).Read()
		m.jmp(uint(o))
	}
}

func (m *Machine) jmpnz(v ...interface{}) {
	if len(v) != 2 {
		log.Panic(ArgCountError{"JMPNZ", 2, uint(len(v))})
	}
	a := v[0].(Reader).Read()
	if a != 0 {
		o := v[1].(Reader).Read()
		m.jmp(uint(o))
	}
}

func (m *Machine) rboff(v ...interface{}) {
	if len(v) != 1 {
		log.Panic(ArgCountError{"RBOFF", 1, uint(len(v))})
	}
	m.rb += v[0].(Reader).Read()
}

func (m *Machine) halt(v ...interface{}) {
	if len(v) != 0 {
		log.Panic(ArgCountError{"HALT", 0, uint(len(v))})
	}
	m.halted = true
}

func (m *Machine) Step() bool {
	if !m.fetch.Fetch() {
		return false
	}
	m.fetch.Instruct().Exec()
	return m.halted
}

func (m *Machine) Run() {
	for m.fetch.Fetch() {
		m.fetch.Instruct().Exec()
	}
}

func (m *Machine) In(in Reader) {
	m.in = in
}

func (m *Machine) Out(out Writer) {
	m.out = out
}

func (m *Machine) Load(p []int) {
	for i, v := range p {
		m.mem.Store(uint(i), v)
	}
}