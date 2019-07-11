package noise

import (
	"bufio"
	"github.com/phyrwork/goadvent/iterator"
	"io"
)

type Reader interface {
	iterator.Iterator
	String() string
}

type Scanner struct {
	*iterator.ScannerIterator
}

func (it *Scanner) String() string { return it.Text() }

func NewScanner(r io.Reader) *Scanner {
	it := &Scanner{iterator.NewScannerIterator(r)}
	it.Split(bufio.ScanLines)
	return it
}

type Slice struct {
	a []string
	i int
}

func NewSlice(a ...string) *Slice {
	return &Slice{a, -1}
}

func (it *Slice) Next() bool {
	it.i++
	return it.i < len(it.a)
}

func (it *Slice) Value() interface{} { return it.a[it.i] }

func (it *Slice) String() string { return it.a[it.i] }

func (it *Slice) Err() error { return nil }
