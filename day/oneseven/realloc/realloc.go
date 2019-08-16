package realloc

import (
	"bufio"
	"io"
	"strconv"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

type Memory []int

type Tree map[int]Tree

func (t Tree) Put(seq ...int) {
	m := t
	for _, c := range seq {
		n := m[c]
		if n == nil {
			n = make(Tree)
			m[c] = n
		}
		m = n
	}
}

func (t Tree) Has(seq ...int) bool {
	m := t
	for _, c := range seq {
		n := m[c]
		if n == nil {
			return false
		}
		m = n
	}
	return true
}

func Pick(m Memory) int {
	i, h := MaxInt, 0
	for j, n := range m {
		if n > h {
			// highest count wins
			i, h = j, n
		} else if n == h {
			// or if tie, lowest index wins
			if j < i {
				i = j
			}
		}
	}
	return i
}

func Distrib(m Memory, i int) {
	var n int
	n, m[i] = m[i], 0
	for n > 0 {
		i = (i + 1) % len(m)
		m[i], n = m[i] + 1, n - 1
	}
}

type Distributor struct {
	Mem    Memory
	hist   Tree
	cycles int
}

func NewDistributor(m Memory) *Distributor {
	d := &Distributor{
		Mem:    m,
		hist:   make(Tree),
		cycles: 0,
	}
	d.hist.Put(m...)
	return d
}

func (d *Distributor) Cycles() int { return d.cycles }

func (d *Distributor) Distrib() {
	for {
		i := Pick(d.Mem)
		Distrib(d.Mem, i)
		d.cycles++
		if d.hist.Has(d.Mem...) {
			break
		}
		d.hist.Put(d.Mem...)
	}
}

func Parse(r io.Reader) (Memory, error) {
	m := make(Memory, 0)
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		s := sc.Text()
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		m = append(m, n)
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	return m, nil
}

func SolveFirst(r io.Reader) (string, error) {
	m, err := Parse(r)
	if err != nil {
		return "", err
	}
	d := NewDistributor(m)
	d.Distrib()
	return strconv.Itoa(d.Cycles()), nil
}

func SolveSize(r io.Reader) (string, error) {
	m, err := Parse(r)
	if err != nil {
		return "", err
	}
	d := NewDistributor(m)
	d.Distrib()
	d = NewDistributor(m)
	d.Distrib()
	return strconv.Itoa(d.Cycles()), nil
}