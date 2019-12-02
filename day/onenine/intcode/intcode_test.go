package intcode

import (
	"testing"
)

func cmpProgram(a, b Program) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestMachine(t *testing.T) {
	m := Machine{
		m:  []int{1,9,10,3,2,3,11,0,99,30,40,50},
		op: DefaultOps,
	}
	for m.Next() {
		// OK
	}
	if err := m.Err(); err != nil {
		t.Fatalf("unexepcted error: %v,", err)
	}
	want := []int{3500,9,10,70,2,3,11,0,99,30,40,50}
	if !cmpProgram(want, Program(m.m)) {
		t.Fatalf("unexpected state: want %v, got %v", want, m.m)
	}
}

