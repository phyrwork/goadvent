package iterator

import "fmt"

type Iterator interface {
	Next() bool
	Value() interface{}
	Err() error
}

type Resetter interface {
	Reset() error
}

type ResetIterator interface {
	Iterator
	Resetter
}

func Each(it Iterator, f func (interface{}) error) error {
	for it.Next() {
		if err := f(it.Value()); err != nil {
			return fmt.Errorf("each func error: %v", err)
		}
	}
	if err := it.Err(); err != nil {
		return fmt.Errorf("iterator error: %v", err)
	}
	return nil
}

func Array(it Iterator) ([]interface{}, error) {
	a := make([]interface{}, 0)
	return a, Each(it, func (v interface{}) error {
		a = append(a, v)
		return nil
	})
}