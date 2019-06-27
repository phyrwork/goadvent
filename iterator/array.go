package iterator

type ArrayIterator struct {
	a []interface{}
	i int
}

func NewArrayIterator(a ...interface{}) *ArrayIterator { return &ArrayIterator{a, -1} }

func (it *ArrayIterator) Reset() error {
	it.i = -1
	return nil
}

func (it *ArrayIterator) Next() bool {
	it.i++
	return it.i < len(it.a)
}

func (it *ArrayIterator) Value() interface{} {
	return it.a[it.i]
}

func (it *ArrayIterator) Err() error { return nil }