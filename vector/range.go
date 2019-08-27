package vector

type Container interface {
	Contains(Vector) bool
}

type Iterator interface {
	Each(func (Vector) error) error
}

type Region interface {
	Container
	Iterator
}

// Range represents a <insert term for 'square in N dimensions' here> region
//
// The start and end vectors are expected to be the same length
//
// s[d] <= e[d] for all dimensions to simplify
type Range struct {
	s, e Vector
}

func NewRange(a, b Vector) Range {
	var l, m Vector // long, min (len)
	if len(a) > len(b) {
		l, m = a, b
	} else {
		l, m = b, a
	}
	s, e := make(Vector, len(l)), make(Vector, len(l)) // start, end
	// Sort dimension values so that s[d] <= e[d]
	for d := 0; d < len(m); d++ {
		if l[d] < m[d] {
			s[d], e[d] = l[d], m[d]
		} else {
			s[d], e[d] = m[d], l[d]
		}
	}
	// Infer zeros where vectors lengths are unequal
	for d := len(m); d < len(l); d++ {
		if l[d] < 0 {
			s[d], e[d] = l[d], 0
		} else {
			s[d], e[d] = 0, l[d]
		}
	}
	return Range{s, e}
}

func (r Range) Start() Vector {
	s := make(Vector, len(r.s))
	copy(s, r.s)
	return s
}

func (r Range) End() Vector {
	e := make(Vector, len(r.e))
	copy(e, r.e)
	return e
}

func (r Range) Contains(a Vector) bool {
	for d := range r.s {
		if a[d] < r.s[d] {
			return false
		}
		if a[d] > r.e[d] {
			return false
		}
	}
	return true
}

func (r Range) Each(f func(Vector) error) error {
	c := make(Vector, len(r.s))
	copy(c, r.s)
	var each func (int) error
	each = func (d int) error {
		for c[d] = r.s[d]; c[d] <= r.e[d]; c[d]++ {
			e := d + 1 // next dimension
			if e < len(c) {
				if err := each(e); err != nil {
					return err
				}
			} else {
				o := make(Vector, len(c))
				copy(o, c)
				if err := f(o); err != nil {
					return err
				}
			}
		}
		return nil
	}
	return each(0)
}