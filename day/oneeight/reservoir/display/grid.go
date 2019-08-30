package display

import "github.com/phyrwork/goadvent/vector"

type Grid [][]rune

func NewGrid(w, h int) Grid {
	g := make([][]rune, w)
	for x := range g {
		g[x] = make([]rune, h)
	}
	return g
}

func (g Grid) Size() vector.Vector {
	w := len(g)
	h := 0
	if w > 0 {
		h = len(g[0]) // assume square
	}
	return vector.Vector{w, h}
}

func (g Grid) Range() *vector.Range {
	w, h := XY(g.Size())
	if w == 0 || h == 0 {
		return nil
	}
	s := vector.Vector{0, 0}
	e := vector.Vector{w - 1, h - 1}
	r := vector.NewRange(s, e)
	return &r
}

func (g Grid) String() string {
	w, h := XY(g.Size())
	s := make([]rune, 0, w*h + h)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			s = append(s, g[x][y])
		}
		s = append(s, '\n')
	}
	return string(s)
}

func (g Grid) Set(v vector.Vector, c rune) error {
	if r := g.Range(); !r.Contains(v) {
		return OutOfBoundsError{*r, v}
	}
	x, y := XY(v)
	g[x][y] = c
	return nil
}

func (g Grid) Get(v vector.Vector) (rune, error) {
	if r := g.Range(); !r.Contains(v) {
		return 0, OutOfBoundsError{*r, v}
	}
	x, y := XY(v)
	c := g[x][y]
	return c, nil
}