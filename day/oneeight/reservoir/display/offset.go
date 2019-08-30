package display

import "github.com/phyrwork/goadvent/vector"

type Offset struct {
	Display
	Offset vector.Vector
}

func (o Offset) Range() *vector.Range {
	r := o.Display.Range()
	if r == nil {
		return nil
	}
	s := r.Offset(o.Offset)
	return &s
}

func (o Offset) Set(v vector.Vector, c rune) error {
	v = vector.Diff(o.Offset, v)
	return o.Display.Set(v, c)
}

func (o Offset) Get(v vector.Vector) (rune, error) {
	v = vector.Diff(o.Offset, v)
	return o.Display.Get(v)
}
