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

func TestMachinePrograms(t *testing.T) {
	tests := map[string]struct {
		in int
		p []int
		want int
	}{
		"equal, pos, false": {
			7,
			[]int{3,9,8,9,10,9,4,9,99,-1,8},
			0,
		},
		"equal, pos, true": {
			8,
			[]int{3,9,8,9,10,9,4,9,99,-1,8},
			1,
		},
		"less, pos, false": {
			7,
			[]int{3,9,7,9,10,9,4,9,99,-1,8},
			1,
		},
		"less, pos, true": {
			8,
			[]int{3,9,7,9,10,9,4,9,99,-1,8},
			0,
		},
		"equal, imm, false": {
			7,
			[]int{3,3,1108,-1,8,3,4,3,99},
			0,
		},
		"equal, imm, true": {
			8,
			[]int{3,3,1108,-1,8,3,4,3,99},
			1,
		},
		"less, imm, false": {
			7,
			[]int{3,3,1107,-1,8,3,4,3,99},
			1,
		},
		"less, imm, true": {
			8,
			[]int{3,3,1107,-1,8,3,4,3,99},
			0,
		},
		"cmp program, lt": {
			7,
			[]int{
				3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,
				1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,
				999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99,
			},
			999,
		},
		"cmp program, eq": {
			8,
			[]int{
				3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,
				1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,
				999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99,
			},
			1000,
		},
		"cmp program, gt": {
			9,
			[]int{
				3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,
				1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,
				999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99,
			},
			1001,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var got int
			out := false
			m := Machine{
				m:  test.p,
				op: DefaultOps,
				in: func () int { return test.in },
				out: func (i int) { out = true; got = i },
			}
			for m.Next() {
				// OK
			}
			if err := m.Err(); err != nil {
				t.Fatalf("unexepcted error: %v,", err)
			}
			if !out {
				t.Fatalf("expected output")
			}
			if got != test.want {
				t.Fatalf("unexpected output: want %v, got %v", test.want, got)
			}
		})
	}
}
