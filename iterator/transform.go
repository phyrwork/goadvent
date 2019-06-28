package iterator

import "fmt"

type TransformIterator struct {
	it   Iterator
	t    func (interface{}) (interface{}, error)
	curr interface{}
	err  error
}

func NewTransformIterator(it Iterator, t func (interface{}) (interface{}, error)) *TransformIterator {
	return &TransformIterator{it, t, nil, nil}
}

func (it *TransformIterator) Reset() error {
	rst, ok := it.it.(ResetIterator)
	if !ok {
		return fmt.Errorf("sub iterator can not reset")
	}
	it.curr = nil
	it.err = rst.Reset()
	return it.err
}

func (it *TransformIterator) Next() bool {
	if !it.it.Next() {
		it.err = it.it.Err()
		return false
	}
	var err error
	it.curr, err = it.t(it.it.Value())
	if err != nil {
		it.err = fmt.Errorf("adapter error: %v", err)
		return false
	}
	return true
}

func (it *TransformIterator) Value() interface{} { return it.curr }

func (it *TransformIterator) Err() error { return it.err }