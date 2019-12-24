package intcode2

import (
	"reflect"
	"testing"
)

func TestMachine(t *testing.T) {
	tests := map[string]struct {
		in int
		p []int
		want []int
	}{
		"equal, pos, false": {
			7,
			[]int{3,9,8,9,10,9,4,9,99,-1,8},
			[]int{0},
		},
		"equal, pos, true": {
			8,
			[]int{3,9,8,9,10,9,4,9,99,-1,8},
			[]int{1},
		},
		"less, pos, true": {
			7,
			[]int{3,9,7,9,10,9,4,9,99,-1,8},
			[]int{1},
		},
		"less, pos, false": {
			8,
			[]int{3,9,7,9,10,9,4,9,99,-1,8},
			[]int{0},
		},
		"equal, imm, false": {
			7,
			[]int{3,3,1108,-1,8,3,4,3,99},
			[]int{0},
		},
		"equal, imm, true": {
			8,
			[]int{3,3,1108,-1,8,3,4,3,99},
			[]int{1},
		},
		"less, imm, true": {
			7,
			[]int{3,3,1107,-1,8,3,4,3,99},
			[]int{1},
		},
		"less, imm, false": {
			8,
			[]int{3,3,1107,-1,8,3,4,3,99},
			[]int{0},
		},
		"cmp program, lt": {
			7,
			[]int{
				3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,
				1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,
				999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99,
			},
			[]int{999},
		},
		"cmp program, eq": {
			8,
			[]int{
				3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,
				1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,
				999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99,
			},
			[]int{1000},
		},
		"cmp program, gt": {
			9,
			[]int{
				3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,
				1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,
				999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99,
			},
			[]int{1001},
		},
		"output self": {
			1,
			[]int{109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99},
			[]int{109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99},
		},
		"output 16 digit": {
			1,
			[]int{1102,34915192,34915192,7,4,7,99,0},
			[]int{34915192 * 34915192},
		},
		"output offset large": {
			1,
			[]int{104,1125899906842624,99},
			[]int{1125899906842624},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := make([]int, 0)
			mem := make(MapMemory)
			m := NewMachine(mem)
			m.In(ReadFunc(func () int {
				return test.in
			}))
			m.Out(WriteFunc(func (v int) {
				got = append(got, v)
			}))
			m.Load(test.p)
			m.Run()
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("unexpected output: want %v, got %v", test.want, got)
			}
		})
	}
}