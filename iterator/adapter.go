package iterator

import "fmt"

type AdapterIterator struct {
	it Iterator
	ad func (interface{}) (interface{}, error)
	curr interface{}
	err error
}

func NewAdapterIterator(it Iterator, ad func (interface{}) (interface{}, error)) *AdapterIterator {
	return &AdapterIterator{it, ad, nil, nil}
}

func (it *AdapterIterator) Reset() error {
	rst, ok := it.it.(ResetIterator)
	if !ok {
		return fmt.Errorf("sub iterator can not reset")
	}
	it.curr = nil
	it.err = rst.Reset()
	return it.err
}

func (it *AdapterIterator) Next() bool {
	if !it.it.Next() {
		it.err = it.it.Err()
		return false
	}
	var err error
	it.curr, err = it.ad(it.it.Value())
	if err != nil {
		it.err = fmt.Errorf("adapter error: %v", err)
		return false
	}
	return true
}

func (it *AdapterIterator) Value() interface{} { return it.curr }

func (it *AdapterIterator) Err() error { return it.err }