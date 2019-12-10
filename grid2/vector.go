package grid2

type Vector [2]int

func (a Vector) Mul(m int) Vector {
	return Vector{a[0]*m, a[1]*m}
}

func (a Vector) Sub(b Vector) Vector {
	return Vector{a[0]-b[0], a[1]-b[1]}
}

