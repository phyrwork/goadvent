package intcode2

import (
	"bufio"
	"fmt"
	"github.com/phyrwork/goadvent/app"
	"github.com/phyrwork/goadvent/iterator"
	"io"
	"strconv"
)

func Read(r io.Reader) ([]int, error) {
	p := make([]int, 0)
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

func Solve5(r io.Reader) app.Solution {
	out := make([]int, 0)
	p, err := Read(r)
	if err != nil {
		return app.Errorf("read error: %v", err)
	}
	mem := make(MapMemory)
	m := NewMachine(mem)
	m.In(ReadFunc(func () int {
		return 1
	}))
	m.Out(WriteFunc(func (i int) {
		out = append(out, i)
	}))
	m.Load(p)
	m.Run()
	if len(out) > 2 {
		return app.Errorf("program error: %v", out)
	}
	return app.Int(out[0])
}