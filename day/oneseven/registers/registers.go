package registers

import (
	"io"
	"strconv"
)

// TODO: there's a bug when reading from stdin where punctuation e.g. ! or = isn't recognised by the parser
// TODO: pasting the puzzle input in as a string literal and running in a test produces the correct answer

func SolveMaxEnd(r io.Reader) (string, error) {
	p, err := Parse(r)
	if err != nil {
		return "", err
	}
	m := Machine{
		Opset: Opset,
		R:     make(Registers),
	}
	m.Exec(p...)
	return strconv.Itoa(m.R.Max()), nil
}

func SolveMaxAny(r io.Reader) (string, error) {
	p, err := Parse(r)
	if err != nil {
		return "", err
	}
	m := Machine{
		Opset: Opset,
		R:     make(Registers),
	}
	v := MinInt
	m.PostStmt = func() {
		w := m.R.Max()
		if w > v {
			v = w
		}
	}
	m.Exec(p...)
	return strconv.Itoa(v), nil
}
