package realloc

import (
	"bufio"
	"github.com/phyrwork/goadvent/iterator"
	"io"
	"strconv"
	"strings"
)

type Memory []int

func NewMemory(m ...int) Memory { return m }

func (m Memory) String() string {
	a := make([]string, len(m))
	for i, n := range m {
		a[i] += strconv.Itoa(n)
	}
	return strings.Join(a, " ")
}

func ScanMemory(r io.Reader) (Memory, error) {
	it := iterator.NewScannerIterator(r)
	it.Split(bufio.ScanWords)
	a := make(Memory, 0)
	return a, iterator.Each(it, func (v interface{}) error {
		s := v.(string)
		n, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		a = append(a, n)
		return nil
	})
}

type MemoryFunc func (m Memory)

// TODO: there's a bug in here somewhere, apparently
// TODO: my part 1 solution is terminating too early
func Realloc(m Memory) {
	// largest
	max := 0
	for b, n := range m {
		if n < m[max] {
			continue
		}
		if n == m[max] && max < n {
			continue
		}
		max = b
	}
	// redist
	n := m[max]
	m[max] = 0
	b := max
	for n > 0 {
		// move
		b++
		if b >= len(m) { // wrap
			b = 0
		}
		// alloc
		m[b], n = m[b] + 1, n - 1
	}
}

type History [][]int

func NewHistory(m ...[]int) *History {
	h := make(History, 0)
	for _, n := range m {
		h.Store(n)
	}
	return &h
}

func (h *History) Store(m Memory) {
	n := make(Memory, len(m))
	copy(n, m)
	*h = append(*h, n)
}

func (h *History) Find(m Memory) bool {
	for _, n := range *h {
		eq := func (a, b Memory) bool {
			if len(n) != len(m) {
				return false
			}
			for i := range a {
				if a[i] != b[i] {
					return false
				}
			}
			return true
		}
		if eq(n, m) {
			return true
		}
	}
	return false
}

func ReallocUniq(m Memory) int {
	h := NewHistory(m)
	c := 0
	for {
		Realloc(m)
		c++
		if h.Find(m) {
			break
		}
		h.Store(m)
	}
	return c
}

func Solve(r io.Reader) (string, error) {
	m, err := ScanMemory(r)
	if err != nil {
		return "", err
	}
	c := ReallocUniq(m)
	return strconv.Itoa(c), nil
}