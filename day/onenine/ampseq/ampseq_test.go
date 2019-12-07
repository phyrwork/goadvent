package ampseq

import (
	"github.com/phyrwork/goadvent/day/onenine/intcode"
	"strconv"
	"testing"
)

func TestSequence(t *testing.T) {
	tests := map[int]struct {
		p intcode.Program
		ph []int
		sig int
	}{
		1: {
			[]int{3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0},
			[]int{4,3,2,1,0},
			43210,
		},
		2: {
			[]int{
				3,23,3,24,1002,24,10,24,1002,23,-1,23,
				101,5,23,23,1,24,23,23,4,23,99,0,0,
			},
			[]int{0,1,2,3,4},
			54321,
		},
		3: {
			[]int{
				3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,
				1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0,
			},
			[]int{1,0,4,3,2},
			65210,
		},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			seq := NewSequence(test.p, 0, test.ph...)
			sig, err := seq.Run()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if sig != test.sig {
				t.Fatalf("unexpected signal: want %v, got %v", test.sig, sig)
			}
		})
	}
}

