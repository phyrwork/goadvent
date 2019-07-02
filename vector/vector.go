package vector

type Vector []int

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
		x[d] -= v[d]
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