package registers

import (
	"github.com/phyrwork/goadvent/app"
	"io"
)

// TODO: there's a bug when reading from stdin where punctuation e.g. ! or = isn't recognised by the parser
// TODO: pasting the puzzle input in as a string literal and running in a test produces the correct answer

func SolveMaxEnd(r io.Reader) app.Solution {
	p, err := Parse(r)
	if err != nil {
		return app.NewError(err)
	}
	m := Machine{
		Opset: Opset,
		R:     make(Registers),
	}
	m.Exec(p...)
	return app.Int(m.R.Max())
}

func SolveMaxAny(r io.Reader) app.Solution {
	p, err := Parse(r)
	if err != nil {
		return app.NewError(err)
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
	return app.Int(v)
}
