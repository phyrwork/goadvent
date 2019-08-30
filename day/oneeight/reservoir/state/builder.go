package state

import "github.com/phyrwork/goadvent/vector"

type MapBuilder struct {
	M Map
}

func (b MapBuilder) Create(r vector.Range, f func () Tile) {
	_ = r.Each(func (v vector.Vector) error {
		t := f()
		b.M.Set(v, t)
		return nil
	})
}
