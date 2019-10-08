package iterator

import "container/ring"

type RingIterator struct {
	r   *ring.Ring // ring
	cur *ring.Ring // cursor
}

func NewRingIterator(r *ring.Ring) *RingIterator {
	return &RingIterator{
		r:   r,
		cur: nil,
	}
}

func (it *RingIterator) Next() bool {
	if it.r == nil {
		// empty ring
		return false
	}
	if it.cur == nil {
		// first call
		it.cur = it.r
		return true
	}
	if it.cur = it.cur.Next(); it.cur == it.r {
		// back to start
		return false
	}
	return true
}

func (it *RingIterator) Value() interface{} {
	if it.cur == nil {
		return nil
	}
	return it.cur.Value
}

func (it *RingIterator) Err() error { return nil }

func (it *RingIterator) Reset() error {
	it.cur = nil
	return nil
}