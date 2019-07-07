package iterator

import (
	"bufio"
	"errors"
	"io"
	"unicode/utf8"
)

func SplitComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// see bufio.ScanWords for implementation notes
	for width, i := 0, 0; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if r == ',' || r == '\n' {
			return i + width, data[0:i], nil
		}
	}
	if atEOF && len(data) > 0 {
		return len(data), data[0:], nil
	}
	return 0, nil, nil
}

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

