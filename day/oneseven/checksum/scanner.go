package checksum

import (
	"bufio"
	"fmt"
	"github.com/phyrwork/goadvent/iterator"
	"io"
	"strconv"
	"strings"
)

type ValueScanner struct {
	sc  *iterator.ScannerIterator
	i   int
	err error
}

func NewValueScanner(r io.Reader) *ValueScanner {
	sc := iterator.NewScannerIterator(r)
	sc.Split(bufio.ScanWords)
	return &ValueScanner{sc, 0, nil}
}

func (it *ValueScanner) Reset() error {
	if err := it.sc.Reset(); err != nil {
		return err
	}
	it.i = 0
	return nil
}

func (it *ValueScanner) Next() bool {
	if it.sc.Scan() {
		s := it.sc.Text()
		var err error
		it.i, err = strconv.Atoi(s)
		if err != nil {
			it.err = fmt.Errorf("error parsing value: %v", err)
			return false
		}
		return true
	}
	if err := it.sc.Err(); err != nil {
		it.err = err
	}
	return false
}

func (it *ValueScanner) Value() interface{} { return it.i }

func (it *ValueScanner) Int() int { return it.i }

func (it *ValueScanner) Err() error { return it.err }

type RowScanner struct {
	sc  *iterator.ScannerIterator
	row Values
	err error
}

func NewRowScanner(r io.Reader) *RowScanner {
	sc := iterator.NewScannerIterator(r)
	sc.Split(bufio.ScanLines)
	return &RowScanner{sc, nil, nil}
}

func (it *RowScanner) Reset() error {
	if err := it.sc.Reset(); err != nil {
		return err
	}
	it.row = nil
	return nil
}

func (it *RowScanner) Next() bool {
	if it.sc.Scan() {
		s := it.sc.Text()
		rd := strings.NewReader(s)
		it.row = NewValueScanner(rd)
		return true
	}
	if err := it.sc.Err(); err != nil {
		it.err = err
	}
	return false
}

func (it *RowScanner) Row() Values { return it.row }

func (it *RowScanner) Err() error { return it.sc.Err() }