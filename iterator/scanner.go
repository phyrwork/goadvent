package iterator

import (
	"bufio"
	"errors"
	"io"
)

type ScannerIterator struct {
	r io.Reader
	*bufio.Scanner
	split bufio.SplitFunc
}

func NewScannerIterator(r io.Reader) *ScannerIterator {
	return &ScannerIterator{r, bufio.NewScanner(r), nil}
}

// i'm sorry :(
func (it *ScannerIterator) Split(f bufio.SplitFunc) {
	it.split = f
	it.Scanner.Split(f)
}

func (it *ScannerIterator) Reset() error {
	sk, ok := it.r.(io.ReadSeeker)
	if !ok {
		return errors.New("not a seeker")
	}
	if _, err := sk.Seek(0, 0); err != nil {
		return err
	}
	it.Scanner = bufio.NewScanner(it.r)
	// yuck :(
	if it.split != nil {
		it.Scanner.Split(it.split)
	}
	return nil
}

func (it *ScannerIterator) Next() bool { return it.Scan() }

func (it *ScannerIterator) Value() interface{} { return it.Text() }