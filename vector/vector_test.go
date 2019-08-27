package vector

import "testing"

var cmpeqTests = map[string]struct{
	a, b Vector
	cmp  bool
	eq   bool
}{
	"len eq, val eq": {Vector{0,1,2}, Vector{0,1,2}, true,  true },
	"len ne, val ne": {Vector{0,1},   Vector{0,1,2}, false, false},
	"len eq, val ne": {Vector{0,1,2}, Vector{0,2,1}, false, false},
	"len ne, val eq": {Vector{0,1},   Vector{0,1,0}, false, true },
}

func TestCmp(t *testing.T) {
	for name, test := range cmpeqTests {
		t.Run(name, func(t *testing.T) {
			cmp := Cmp(test.a, test.b)
			if cmp != test.cmp {
				t.Fatalf("unexpected result: want %v, got %v", test.cmp, cmp)
			}
		})
	}
}

func TestEq(t *testing.T) {
	for name, test := range cmpeqTests {
		t.Run(name, func(t *testing.T) {
			eq := Eq(test.a, test.b)
			if eq != test.eq {
				t.Fatalf("unexpected result: want %v, got %v", test.eq, eq)
			}
		})
	}
}



