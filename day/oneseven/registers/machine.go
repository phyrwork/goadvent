package registers

type Registers map[string]int

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func (r Registers) Max() int {
	m := MinInt
	for _, v := range r {
		if v > m {
			m = v
		}
	}
	return m
}

type OpFn func (*int, int) int

func inc(r *int, a int) int {
	*r += a
	return 0
}

func dec(r *int, a int) int {
	*r -= a
	return 0
}

func eq(r *int, a int) int {
	if *r == a {
		return 1
	} else {
		return 0
	}
}
func neq(r *int, a int) int {
	if *r != a {
		return 1
	} else {
		return 0
	}
}

func gt(r *int, a int) int {
	if *r > a {
		return 1
	} else {
		return 0
	}
}

func gte(r *int, a int) int {
	if *r >= a {
		return 1
	} else {
		return 0
	}
}

func lt(r *int, a int) int {
	if *r < a {
		return 1
	} else {
		return 0
	}
}

func lte(r *int, a int) int {
	if *r <= a {
		return 1
	} else {
		return 0
	}
}

var Opset = map[string]OpFn{
	OpInc: inc,
	OpDec: dec,
	OpEq: eq,
	OpNeq: neq,
	OpGt: gt,
	OpGte: gte,
	OpLt: lt,
	OpLte: lte,
}

type Machine struct {
	Opset map[string]OpFn
	R Registers
	PostStmt func ()
}

// Make it work and then make it good?
// All unset registers default to 0 (as per spec) thanks to Go maps
func (m *Machine) Exec(stmts ...Stmt) {
	for _, stmt := range stmts {
		r, cmp, a := m.R[stmt.Cond.Reg], m.Opset[stmt.Cond.Op], stmt.Cond.Arg
		if cmp(&r, a) != 0 {
			r, mod, a := m.R[stmt.Op.Reg], m.Opset[stmt.Op.Op], stmt.Op.Arg
			mod(&r, a)
			m.R[stmt.Op.Reg] = r
		}
		if m.PostStmt != nil {
			m.PostStmt()
		}
	}
}