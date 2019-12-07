package intcode

import (
	"bufio"
	"fmt"
	"github.com/phyrwork/goadvent/app"
	"github.com/phyrwork/goadvent/iterator"
	"io"
	"strconv"
)

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

func Solve3(r io.Reader) app.Solution {
	out := make([]int, 0)
	p, err := Read(r)
	if err != nil {
		return app.Errorf("read error: %v", err)
	}
	m := Machine{
		m:    Memory(p),
		op:   DefaultOps,
		Read: func () int { return 1 },
		Write: func (i int) {
			out = append(out, i)
		},
	}
	for m.Next() {}
	if err := m.Err(); err != nil {
		return app.Errorf("run error: %v", err)
	}
	if len(out) == 0 {
		return app.Errorf("program produced no output")
	}
	return app.Int(out[len(out)-1])
}

func Solve4(r io.Reader) app.Solution {
	out := make([]int, 0)
	p, err := Read(r)
	if err != nil {
		return app.Errorf("read error: %v", err)
	}
	m := Machine{
		m:    Memory(p),
		op:   DefaultOps,
		Read: func () int { return 5 },
		Write: func (i int) {
			out = append(out, i)
		},
	}
	for m.Next() {}
	if err := m.Err(); err != nil {
		return app.Errorf("run error: %v", err)
	}
	if len(out) == 0 {
		return app.Errorf("program produced no output")
	}
	return app.Int(out[0])
}