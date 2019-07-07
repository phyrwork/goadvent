package knot

// TODO: whilst this was a fun way of playing about with concurrency,
//  (hello errgroup) by the end of part 2 it is clear that there
//  there is no real need for channels here

import (
	"bufio"
	"container/ring"
	"encoding/hex"
	"fmt"
	"github.com/phyrwork/goadvent/iterator"
	"golang.org/x/sync/errgroup"
	"io"
	"strconv"
)

func NewRing(n int) *ring.Ring {
	if n > 256 {
		panic("n > 256")
	}
	r := ring.New(int(n))
	for i := 0; i < n; i++ {
		r.Value = byte(i)
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

func SparseHash(n int, str <-chan byte) []byte {
	r := NewRing(n)
	pos := 0
	skip := 0
	rv := NewReverser(r)
	for size := range str {
		// Although Reverser is able to reverse a sector by offset,
		// we can reduce accumulating large offset traversals at each
		// operation by moving the cursor and correcting at the end
		rv.Sector(0, int(size))
		adv := int(size) + skip
		rv.Move(adv)
		pos += adv
		skip++
	}
	h := make([]byte, 0, n)
	r = rv.Move(-pos % rv.Len()).Ring()
	r.Do(func (v interface{}) {
		h = append(h, v.(byte))
	})
	return h
}

func DenseHash(in []byte) ([]byte, error) {
	if len(in) % 16 != 0 {
		return nil, fmt.Errorf("input size (%v) not multiple of block size (%v)", len(in), 16)
	}
	h := make([]byte, 0, len(in)/16)
	for o := 0; o < len(in); o += 16 {
		x := byte(0)
		for n := 0; n < 16; n++ {
			x ^= in[o + n]
		}
		h = append(h, x)
	}
	return h, nil
}

// TODO: there's a bug somewhere (my input doesn't give the right answer)
//  but it passes for all the given examples! :(
func KnotHash(rd io.Reader) (string, error) {
	str := make(chan byte)
	g := errgroup.Group{}
	g.Go(func() error {
		return NewRepeatStream(NewByteStream(rd), 64)(str)
	})
	s := SparseHash(256, str)
	if err := g.Wait(); err != nil {
		return "", err
	}
	d, err := DenseHash(s)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(d), nil
}

type StreamFunc func (str chan <-byte) error

func NewCommaStream(rd io.Reader) StreamFunc {
	return func(out chan <-byte) error {
		defer close(out)
		sc := bufio.NewScanner(rd)
		sc.Split(iterator.SplitComma)
		for sc.Scan() {
			s := sc.Text()
			size, err := strconv.Atoi(s)
			if err != nil {
				return err
			}
			out <- byte(size)
		}
		return sc.Err()
	}
}

func NewByteStream(rd io.Reader) StreamFunc {
	return func(out chan <-byte) error {
		defer close(out)
		sc := bufio.NewScanner(rd)
		sc.Split(bufio.ScanRunes)
		for sc.Scan() {
			s := sc.Text()
			out <- byte(s[0])
		}
		// n.b. although it's not in the naming, just put the weird
		// additional sizes demanded by the part two solution here
		for _, size := range []byte{17, 31, 73, 47, 23} {
			out <- size
		}
		return sc.Err()
	}
}

func NewRepeatStream(in StreamFunc, n int) StreamFunc {
	return func (out chan <-byte) error {
		c := make(chan byte)
		g := errgroup.Group{}
		g.Go(func() error {
			return in(c)
		})
		a := make([]byte, 0)
		for e := range c {
			a = append(a, e)
		}
		if err := g.Wait(); err != nil {
			return fmt.Errorf("stream repeat error: %v", err)
		}
		go func () {
			defer close(out)
			for ; n > 0; n-- {
				for _, i := range a {
					out <- i
				}
			}
		}()
		return nil
	}
}

func solveSparse(n int, strf StreamFunc) (int, error) {
	str := make(chan byte)
	g := errgroup.Group{}
	g.Go(func() error {
		return strf(str)
	})
	h := SparseHash(n, str)
	if err := g.Wait(); err != nil {
		return 0, err
	}
	a, b := h[0], h[1]
	return int(a) * int(b), nil
}

func SolveSparse(rd io.Reader) (string, error) {
	h, err := solveSparse(256, NewCommaStream(rd))
	if err != nil {
		return "", err
	}
	s := strconv.Itoa(h)
	return s, nil
}