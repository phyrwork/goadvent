package knot

import (
	"bufio"
	"container/ring"
	"github.com/phyrwork/goadvent/app"
	"github.com/phyrwork/goadvent/iterator"
	"golang.org/x/sync/errgroup"
	"io"
	"strconv"
)

func New(n int) *ring.Ring {
	r := ring.New(n)
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}
	return r
}

type Reverser struct {
	r *ring.Ring
	len int
}

func NewReverser(r *ring.Ring) *Reverser { return &Reverser{r, r.Len()} }

func (r *Reverser) Ring() *ring.Ring { return r.r }

func (r *Reverser) Len() int { return r.len }

func (r *Reverser) Move(n int) *Reverser {
	r.r = r.r.Move(n)
	return r

}

func (r *Reverser) All() *Reverser {
	p := r.r.Prev() // ring to unlink, such that Next() is first element in r
	s := p.Unlink(1) // reversed ring
	for n := p.Len(); n > 0; n-- {
		u := p.Unlink(1)
		u.Link(s)
		s = u
	}
	r.r = s
	return r
}

func (r *Reverser) Sector(o int, n int) *Reverser {
	if n == 0 {
		return r
	}
	if n > r.len {
		n = r.len
	}
	p := r.r.Move(o - 1)
	s := p.Unlink(n)
	rv := NewReverser(s)
	s = rv.All().Ring()
	if n == r.len {
		r.r = s
		return r
	}
	p.Link(s)
	r.r = p.Move(1 - o)
	return r
}

func Knot(r *ring.Ring, str <-chan int) *ring.Ring {
	pos := 0
	skip := 0
	rv := NewReverser(r)
	for size := range str {
		// Although Reverser is able to reverse a sector by offset,
		// we can reduce accumulating large offset traversals at each
		// operation by moving the cursor and correcting at the end
		rv.Sector(0, size)
		adv := size + skip
		rv.Move(adv)
		pos += adv
		skip++
	}
	// TODO: I have no idea why we need to return Prev() but, yolo, it works
	return rv.Move(-pos % rv.Len()).Ring()
}

func Solve(n int, rd io.Reader) (int, error) {
	sc := bufio.NewScanner(rd)
	sc.Split(iterator.ScanComma)
	str := make(chan int)
	g := errgroup.Group{}
	g.Go(func() error {
		defer close(str)
		for sc.Scan() {
			s := sc.Text()
			size, err := strconv.Atoi(s)
			if err != nil {
				return err
			}
			str <- size
		}
		return sc.Err()
	})
	r := New(n)
	r = Knot(r, str)
	if err := g.Wait(); err != nil {
		return 0, err
	}
	a, b := r.Value.(int), r.Next().Value.(int)
	return a * b, nil
}

func NewSolver() app.SolverFunc {
	return func (rd io.Reader) (string, error) {
		h, err := Solve(256, rd)
		if err != nil {
			return "", err
		}
		s := strconv.Itoa(h)
		return s, nil
	}
}