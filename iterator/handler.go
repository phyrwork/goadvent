package iterator

import "fmt"

type HandlerIterator struct {
	it  Iterator
	f   func (interface{}) error
	err error
}

func NewHandlerIterator(it Iterator, f func (interface{}) error) *HandlerIterator {
	return &HandlerIterator{it, f, nil}
}

func (it *HandlerIterator) Reset() error {
	rst, ok := it.it.(ResetIterator)
	if !ok {
		return fmt.Errorf("sub iterator can not reset")
	}
	it.err = rst.Reset()
	return it.err
}

func (it *HandlerIterator) Next() bool {
	if !it.it.Next() {
		it.err = it.it.Err()
		return false
	}
	err := it.f(it.it.Value())
	if err != nil {
		it.err = fmt.Errorf("adapter error: %v", err)
		return false
	}
	return true
}

func (it *HandlerIterator) Value() interface{} { return nil }

func (it *HandlerIterator) Err() error { return it.err }