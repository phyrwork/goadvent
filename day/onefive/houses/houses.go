package houses

import (
	"github.com/phyrwork/goadvent/app"
	"github.com/phyrwork/goadvent/iterator"
	"io"
	"log"
	"strconv"
)

type coord struct {
	x int
	y int
}

func CountUnique(it iterator.RuneIterator) (int, error) {
	m := make(map[coord]int)
	p := coord{0, 0}
	m[p] = 1
	for it.Next() {
		r := it.Rune()
		switch r {
		case '^':
			p.y++
		case 'v':
			p.y--
		case '>':
			p.x++
		case '<':
			p.x--
		default:
			log.Panicf("unexpected direction: %c", r)
		}
		m[p] = m[p] + 1
	}
	return len(m), it.Err()
}

type RouteFunc func (iterator.RuneIterator) (int, error)

func NewSolver(f RouteFunc) app.SolverFunc {
	return func (r io.Reader) (string, error) {
		sc := iterator.NewRuneScanner(r)
		sc.Skip = iterator.SkipWhitespace
		i, err := f(sc)
		if err != nil {
			return "", err
		}
		return strconv.Itoa(i), nil
	}
}