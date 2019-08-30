package vector

type Vector []int

// Cmp returns true if the given vectors have equal length and value
func Cmp(a, b Vector) bool {
	if len(a) != len(b) {
		return false
	}
	for d := range a {
		if a[d] != b[d] {
			return false
		}
	}
	return true
}

// Eq returns true if the given vectors have equal value
func Eq(a, b Vector) bool {
	var l, m Vector
	if len(a) > len(b) {
		l, m = a, b
	} else {
		l, m = b, a
	}
	for d := 0; d < len(m); d++ {
		if l[d] != m[d] {
			return false
		}
	}
	for d := len(m); d < len(l); d++ {
		if l[d] != 0 {
			return false
		}
	}
	return true
}

func Abs(v Vector) Vector {
	x := make(Vector, len(v))
	copy(x, v)
	for i := range x {
		if x[i] < 0 {
			x[i] = -x[i]
		}
	}
	return x
}

func Sum(v ...Vector) Vector {
	d := 0
	for _, v := range v {
		if e := len(v); e > d {
			d = e
		}
	}
	x := make(Vector, d)
	for _, v := range v {
		for d := range v {
			x[d] += v[d]
		}
	}
	return x
}

func Diff(v, w Vector) Vector {
	d := len(w)
	if e := len(v); e > d {
		d = e
	}
	x := make(Vector, d)
	copy(x, w)
	for d := range v {
		// w - v
		x[d] -= v[d]
	}
	return x
}

func Mult(v Vector, n int) Vector {
	x := make(Vector, len(v))
	for d := range v {
		x[d] = v[d] * n
	}
	return x
}

func Manhattan(v, w Vector) int {
	a := Abs(Diff(v, w))
	s := 0
	for _, m := range a {
		s += m
	}
	return s
}