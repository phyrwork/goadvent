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

func TestMachineBasic(t *testing.T) {
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

func TestMachineReadWrite(t *testing.T) {
	want := 1337
	var got int
	m := Machine{
		m:  []int{3,0,4,0,99},
		op: DefaultOps,
		in: func () int { return want },
		out: func (i int) { got = i },
	}
	for m.Next() {
		// OK
	}
	if err := m.Err(); err != nil {
		t.Fatalf("unexepcted error: %v,", err)
	}
	if got != want {
		t.Fatalf("unexpected output: want %v, got %v", want, got)
	}
}

func TestMachineArgMode(t *testing.T) {
	m := Machine{
		m:  []int{1002,4,3,4,33},
		op: DefaultOps,
	}
	for m.Next() {
		// OK
	}
	if err := m.Err(); err != nil {
		t.Fatalf("unexepcted error: %v,", err)
	}
	if m.m[4] != 99 {
		t.Fatalf("unexpected value at address %v: want %v, got %v", 4, 99, m.m[4])
	}
}

