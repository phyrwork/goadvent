package realloc

import (
	"reflect"
	"strconv"
	"testing"
)

func (m Memory) String() string {
	s := ""
	if len(m) < 1 {
		return s
	}
	s += strconv.Itoa(m[0])
	if len(m) < 2 {
		return s
	}
	for _, n := range m[1:] {
		s += "," + strconv.Itoa(n)
	}
	return s
}

var example = []struct {
	mem     Memory
	pick    int
	distrib Memory
}{
	{Memory{0, 2, 7, 0}, 2, Memory{2, 4, 1, 2}},
	{Memory{2, 4, 1, 2}, 1, Memory{3, 1, 2, 3}},
	{Memory{3, 1, 2, 3}, 0, Memory{0, 2, 3, 4}},
	{Memory{0, 2, 3, 4}, 3, Memory{1, 3, 4, 1}},
	{Memory{1, 3, 4, 1}, 2, Memory{2, 4, 1, 2}},
}

func clone(m Memory) Memory {
	n := make(Memory, len(m))
	copy(n, m)
	return n
}

func TestPick(t *testing.T) {
	for _, test := range example {
		t.Run(test.mem.String(), func(t *testing.T) {
			pick := Pick(test.mem)
			if pick != test.pick {
				t.Fatalf("unexpected selection: want %v, got %v", test.pick, pick)
			}
		})
	}
}

func TestDistrib(t *testing.T) {
	for _, test := range example {
		t.Run(test.mem.String(), func(t *testing.T) {
			mem := clone(test.mem)
			Distrib(mem, test.pick)
			if !reflect.DeepEqual(mem, test.distrib) {
				t.Fatalf("unexpected values: want %v, got %v", test.distrib, mem)
			}
		})
	}
}

func TestTree(t *testing.T) {
	tree := make(Tree)
	tree.Put(2, 4, 1, 2)
	tests := []struct {
		seq  []int
		want bool
	}{
		{[]int{3, 1, 2, 3}, false},
		{[]int{0, 2, 3, 4}, false},
		{[]int{1, 3, 4, 1}, false},
		{[]int{2, 4, 1, 2}, true},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := tree.Has(test.seq...)
			if got != test.want {
				t.Fatalf("unexpected result: want %v, got %v", test.want, got)
			}
		})
	}
}

func TestDistributor(t *testing.T) {
	d := NewDistributor(clone(example[0].mem))
	d.Distrib()
	got := d.Cycles()
	want := len(example)
	if got != want {
		t.Fatalf("unexpected cycle count: want %v, got %v", want, got)
	}
}
