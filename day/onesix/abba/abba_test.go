package abba

import (
	"fmt"
	"strings"
	"testing"
)

type addressOrder []Address

func (l addressOrder) String() string {
	s := ""
	i := 0
	for ; i < len(l) - 1; i++ {
		s += string(l[i]) + "\n"
	}
	s += string(l[i])
	return s
}

func (s Sequence) Eq(t Sequence) bool {
	if len(s) != len(t) {
		return false
	}
	for i := range s {
		if s[i] != t[i] {
			return false
		}
	}
	return true
}

// note to self: not sure if I'm a fan of map[struct(in)]struct(out) test definitions
type sequenceTest struct {
	addr Address
	len  int
}

type sequenceItem struct{
	seq   Sequence
	hyper bool
}

type sequenceOrder []sequenceItem

func (o sequenceOrder) Eq(p sequenceOrder) bool {
	if len(o) != len(p) {
		return false
	}
	for i := range o {
		a, b := o[i], p[i]
		if !a.seq.Eq(b.seq) {
			return false
		}
		if a.hyper != b.hyper {
			return false
		}
	}
	return true
}

func TestSequencer(t *testing.T) {
	tests := map[sequenceTest]sequenceOrder {
		{
			"abba[mnop]qrst",
			4,
		}: {
			{Sequence{'a', 'b', 'b', 'a'}, false},
			{Sequence{'m', 'n', 'o', 'p'}, true},
			{Sequence{'q', 'r', 's', 't'}, false},
		},
		{
			"abba[mnop]qrst",
			3,
		}: {
			{Sequence{'a', 'b', 'b'}, false},
			{Sequence{'b', 'b', 'a'}, false},
			{Sequence{'m', 'n', 'o'}, true},
			{Sequence{'n', 'o', 'p'}, true},
			{Sequence{'q', 'r', 's'}, false},
			{Sequence{'r', 's', 't'}, false},
		},
		{
			"ioxxoj[asdfgh]zxcvbn",
			4,
		}: {
			{Sequence{'i', 'o', 'x', 'x'}, false},
			{Sequence{'o', 'x', 'x', 'o'}, false},
			{Sequence{'x', 'x', 'o', 'j'}, false},
			{Sequence{'a', 's', 'd', 'f'}, true},
			{Sequence{'s', 'd', 'f', 'g'}, true},
			{Sequence{'d', 'f', 'g', 'h'}, true},
			{Sequence{'z', 'x', 'c', 'v'}, false},
			{Sequence{'x', 'c', 'v', 'b'}, false},
			{Sequence{'c', 'v', 'b', 'n'}, false},
		},
	}
	for test, want := range tests {
		t.Run(fmt.Sprintf("%v:%v", test.addr, test.len), func(t *testing.T) {
			got := make(sequenceOrder, 0)
			it := NewSequencer(test.addr, test.len) // TODO: var len tests
			for it.Next() {
				seq := it.Sequence()
				hyper := it.Hyper()
				got = append(got, sequenceItem{seq, hyper})
			}
			if !want.Eq(got) {
				t.Fatalf("unexpected sequences: want %v, got %v", want, got)
			}
		})
	}
}

func TestSupportsTLS(t *testing.T) {
	tests := map[Address]bool {
		"abba[mnop]qrst": true,
		"abcd[bddb]xyyx": false,
		"aaaa[qwer]tyui": false,
		"ioxxoj[asdfgh]zxcvbn": true,
	}
	for addr, want := range tests {
		t.Run(string(addr), func(t *testing.T) {
			got := addr.SupportsTLS()
			if want != got {
				t.Fatalf("unexpected result: want %v, got %v", want, got)
			}
		})
	}
}

func TestSupportsSSL(t *testing.T) {
	tests := map[Address]bool {
		"aba[bab]xyz": true,
		"xyx[xyx]xyx": false,
		"aaa[kek]eke": true,
		"zazbz[bzb]cdb": true,
	}
	for addr, want := range tests {
		t.Run(string(addr), func(t *testing.T) {
			got := addr.SupportsSSL()
			if want != got {
				t.Fatalf("unexpected result: want %v, got %v", want, got)
			}
		})
	}
}

func TestSolver(t *testing.T) {
	tests := map[string]struct {
		list []Address
		filt AddressFilter
		want string
	}{
		"part 1 example": {
			[]Address{
				"abba[mnop]qrst",
				"abcd[bddb]xyyx",
				"aaaa[qwer]tyui",
				"ioxxoj[asdfgh]zxcvbn",
			},
			Address.SupportsTLS,
			"2",
		},
		"part 2 example": {
			[]Address{
				"aba[bab]xyz",
				"xyx[xyx]xyx",
				"aaa[kek]eke",
				"zazbz[bzb]cdb",
			},
			Address.SupportsSSL,
			"3",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			in := strings.NewReader(addressOrder(test.list).String())
			solver := NewSolver(test.filt)
			got, err := solver.Solve(in)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != test.want {
				t.Fatalf("unexpected answer: want %v, got %v", test.want, got)
			}
		})
	}
}

