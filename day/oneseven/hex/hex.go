package hex

import (
	"fmt"
	"github.com/phyrwork/goadvent/app"
	"github.com/phyrwork/goadvent/iterator"
	"github.com/phyrwork/goadvent/vector"
	"io"
	"log"
	"strconv"
)

// https://www.redblobgames.com/grids/hexagons/#distances
// is a very useful resource on hexagonal grids
// tl;dr:
// - every step on the hex grid is two steps in cube space
// - step distance is half the cube manhattan distance
// TODO: generalize this into a 'mesh/hex' package
//
// Arbitarily assign axes to side pairs:
//
//     +x      +z
//       \ n  /
//     nw +--+ ne
//       /    \
// +y --+      +-- -y
//       \    /
//     sw +--+ se
//       / s  \
//     -z      -x
//
var dirs = map[string]vector.Vector {
	//      x   y   z
	"n":  {+1,  0, +1},
	"s":  {-1,  0, -1},
	"ne": { 0, -1, +1},
	"sw": { 0, +1, -1},
	"nw": {+1, +1,  0},
	"se": {-1, -1,  0},
}



func NewScanner(r io.Reader) *iterator.TransformIterator {
	sc := iterator.NewScannerIterator(r)
	sc.Split(iterator.ScanComma)
	return iterator.NewTransformIterator(sc, func (v interface{}) (interface{}, error) {
		d, ok := dirs[v.(string)]
		if !ok {
			return nil, fmt.Errorf("unknown token: %v", d)
		}
		return d, nil
	})
}

func ScanVector(r io.Reader) (vector.Vector, error) {
	d := vector.Vector{0, 0, 0}
	sc := NewScanner(r)
	return d, iterator.Each(sc, func (v interface{}) error {
		d = vector.Sum(d, v.(vector.Vector))
		return nil
	})
}

func Dist(a, b vector.Vector) int {
	if d := len(a); d != 3 {
		log.Panicf("vector a not has len=3: %#v", a)
	}
	if d := len(b); d != 3 {
		log.Panicf("vector b not has len=3: %#v", b)
	}
	return vector.Manhattan(a, b)/2
}

type AccumulatorFunc func (it iterator.Iterator) (vector.Vector, error)

func Sum(it iterator.Iterator) (vector.Vector, error) {
	d := vector.Vector{0, 0, 0}
	return d, iterator.Each(it, func (v interface{}) error {
		d = vector.Sum(d, v.(vector.Vector))
		return nil
	})
}

func MaxDist(it iterator.Iterator) (vector.Vector, error) {
	o := vector.Vector{0, 0, 0}
	d := o
	m := d
	return m, iterator.Each(it, func(v interface{}) error {
		d = vector.Sum(d, v.(vector.Vector))
		if b, a := Dist(o, m), Dist(o, d); a > b {
			m = d
		}
		return nil
	})
}

func Solve(r io.Reader, f AccumulatorFunc) (int, error) {
	sc := NewScanner(r)
	v, err := f(sc)
	if err != nil {
		return 0, err
	}
	return Dist(vector.Vector{0, 0, 0}, v), nil
}

func NewSolver(f AccumulatorFunc) app.SolverFunc {
	return func (r io.Reader) (string, error) {
		d, err := Solve(r, f)
		if err != nil {
			return "", err
		}
		return strconv.Itoa(d), nil
	}
}