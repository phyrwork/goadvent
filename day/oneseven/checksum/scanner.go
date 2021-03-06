package checksum

import (
	"bufio"
	"github.com/phyrwork/goadvent/iterator"
	"io"
	"strconv"
	"strings"
)

func NewValueScanner(r io.Reader) *iterator.TransformIterator {
	it := iterator.NewScannerIterator(r)
	it.Split(bufio.ScanWords)
	return iterator.NewTransformIterator(it, func (v interface{}) (interface{}, error) {
		s := v.(string)
		return strconv.Atoi(s)
	})
}

func NewRowScanner(r io.Reader) *iterator.TransformIterator {
	it := iterator.NewScannerIterator(r)
	it.Split(bufio.ScanLines)
	return iterator.NewTransformIterator(it, func (v interface{}) (interface{}, error) {
		s := v.(string)
		r := strings.NewReader(s)
		return NewValueScanner(r), nil
	})
}