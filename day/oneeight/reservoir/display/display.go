package display

import "github.com/phyrwork/goadvent/vector"

const (
	MarkerSand   = '.'
	MarkerClay   = '#'
	MarkerSource = '+'
	MarkerFlow   = '|'
	MarkerStill  = '~'
)

func XY(v vector.Vector) (int, int) {
	var x, y int
	if len(v) > 0 {
		x = v[0]
	}
	if len(v) > 1 {
		y = v[0]
	}
	return x, y
}

type Display interface {
	Set(v vector.Vector, c rune) error
	Get(v vector.Vector) (rune, error)
	Size() vector.Vector
	Range() *vector.Range
	String() string
}