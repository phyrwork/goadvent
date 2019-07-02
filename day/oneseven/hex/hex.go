package hex

import (
	"bufio"
	"fmt"
	"github.com/phyrwork/goadvent/app"
	"github.com/phyrwork/goadvent/vector"
	"io"
	"log"
	"strconv"
	"unicode/utf8"
)

const toksep = ','

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

func ScanTokens(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// see bufio.ScanWords for implementation notes
	for width, i := 0, 0; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if r == toksep || r == '\n' {
			return i + width, data[0:i], nil
		}
	}
	if atEOF && len(data) > 0 {
		return len(data), data[0:], nil
	}
	return 0, nil, nil
}

func NewScanner(r io.Reader) *bufio.Scanner {
	sc := bufio.NewScanner(r)
	sc.Split(ScanTokens)
	return sc
}

func ScanVector(r io.Reader) (vector.Vector, error) {
	v := vector.Vector{0, 0, 0}
	sc := NewScanner(r)
	for sc.Scan() {
		t := sc.Text()
		d, ok := dirs[t]
		if !ok {
			return nil, fmt.Errorf("unknown token: %v", t)
		}
		v = vector.Sum(v, d)
	}
	return v, nil
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

func Solve(r io.Reader) (int, error) {
	v, err := ScanVector(r)
	if err != nil {
		return 0, nil
	}
	return Dist(vector.Vector{0, 0, 0}, v), nil
}

func NewSolver() app.SolverFunc {
	return func (r io.Reader) (string, error) {
		d, err := Solve(r)
		if err != nil {
			return "", err
		}
		return strconv.Itoa(d), nil
	}
}