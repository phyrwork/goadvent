package houses

import (
	"fmt"
	"github.com/phyrwork/goadvent/iterator"
	"golang.org/x/sync/errgroup"
	"sync"
)

const (
	North = '^'
	South = 'v'
	East  = '>'
	West  = '<'
)

type Direction rune

type DirectionsFunc func (<-chan Direction) error

type Coord struct {
	x, y int
}

type Follow struct {
	p Coord
	Visit func (Coord) error
}

func NewFollow(f Coord) *Follow { return &Follow{f, nil} }

func (f *Follow) Step(d Direction) error {
	switch d {
	case North:
		f.p.y++
	case South:
		f.p.y--
	case East:
		f.p.x++
	case West:
		f.p.x--
	default:
		return fmt.Errorf("unsupported direction: %v", d)
	}
	if f.Visit != nil {
		return f.Visit(f.p)
	}
	return nil
}

func (f *Follow) Each(i <-chan Direction) error {
	for d := range i {
		if err := f.Step(d); err != nil {
			return err
		}
	}
	return nil
}

func Round(f func (int) DirectionsFunc, n int, b int) DirectionsFunc {
	return func (i <-chan Direction) error {
		g := errgroup.Group{}
		ch := make([]chan Direction, n)
		for j := range ch {
			ch[j] = make(chan Direction, b)
			g.Go(func () error {
				return f(n)(ch[j])
			})
		}
		k := 0
		for d := range i {
			ch[k%n] <- d
		}
		return g.Wait()
	}
}

type Visited struct {
	mu sync.Mutex
	m  map[Coord]int
}

func (v *Visited) Visit(p Coord) error {
	v.mu.Lock()
	v.m[p] = v.m[p] + 1
	v.mu.Unlock()
	return nil
}

func FromOrigin(r chan <-Coord) DirectionsFunc {
	f := NewFollow(Coord{0, 0})
	f.Visit = func (p Coord) error {
		r <- p
		return nil
	}
	return f.Each
}

func FromOriginRound(r chan <-Coord, n int, b int) DirectionsFunc {
	g := func (int) DirectionsFunc {
		return FromOrigin(r)
	}
	return Round(g, n, b)
}

func Solve(it iterator.RuneIterator) (int, error) {
	i := make(chan Direction)
	g := errgroup.Group{}
	g.Go(func () error {
		defer close(i)
		for it.Next() {
			i <- Direction(it.Rune())
		}
		return it.Err()
	})
	r := make(chan Coord)
	f := FromOrigin(r)
	if err := f(i); err != nil {
		return 0, err
	}
	m := make(map[Coord]int)
	for p := range r {
		m[p] = m[p] + 1
	}
	return len(m), g.Wait()
}