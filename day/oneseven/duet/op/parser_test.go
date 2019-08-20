package op

import (
	"reflect"
	"strings"
	"testing"
)

var example =
`set a 1
add a 2
mul a a
mod a 5
snd a
set a 0
rcv a
jgz a -1
set a 1
jgz a -2`

func newIntValue(i int) Value {
	s := []int{i} // Allocate an int on the heap
	return Value{Integer: &s[0]}
}

func newVarValue(v string) Value {
	s := []string{v}
	return Value{Variable: &s[0]}
}

func cmpValue(a, b Value) bool {
	switch {
	case a.Integer != nil:
		switch {
		case b.Integer == nil:
			return false
		case *a.Integer != *b.Integer:
			return false
		default:
			return true
		}
	case a.Variable != nil:
		switch {
		case b.Variable == nil:
			return false
		case *a.Variable != *b.Variable:
			return false
		default:
			return true
		}
	default:
		panic("value is neither locator nor integer")
	}
}

func TestParse(t *testing.T) {
	tests := map[string]struct {
		in   string
		want []Op
	}{
		"set l v": {
			"set a 1",
			[]Op{
				{Set: &Set{L:"a", V: newIntValue(1)}},
			},
		},
		"add l v": {
			"add a 2",
			[]Op{
				{Add: &Add{L:"a", V: newIntValue(2)}},
			},
		},
		"mul l l": {
			"mul a a",
			[]Op{
				{Mul: &Mul{L:"a", V: newVarValue("a")}},
			},
		},
		"jgz l v": {
			"jgz a -1",
			[]Op{
				{Jgz: &Jgz{C: newVarValue("a"), O: newIntValue(-1)}},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := Parse(strings.NewReader(test.in))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(got) != len(test.want) {
				t.Fatalf("unexpected count: want %v, got %v", len(test.want), len(got))
			}
			for n := range got {
				ig, iw := got[n].Op(), test.want[n].Op()
				if tg, tw := reflect.TypeOf(ig), reflect.TypeOf(iw); tg != tw {
					t.Fatalf("unexpected type: want %v, got %v", tw, tg)
				}
				switch g := ig.(type) {
				// TODO: these tests feel a bit messy
				case Snd:
					w := iw.(Snd)
					if !cmpValue(w.V, g.V) {
						// TODO: may have to write a String() for Value to deref. pointers
						t.Fatalf("unexpected value: want %v, got %v", w.V, g.V)
					}
				case Set:
					w := iw.(Set)
					if w.L != g.L {
						t.Fatalf("unexpected locator: want %v, got %v", w.L, g.L)
					}
					if !cmpValue(w.V, g.V) {
						t.Fatalf("unexpected value: want %v, got %v", w.V, g.V)
					}
				case Add:
					w := iw.(Add)
					if w.L != g.L {
						t.Fatalf("unexpected locator: want %v, got %v", w.L, g.L)
					}
					if !cmpValue(w.V, g.V) {
						t.Fatalf("unexpected value: want %v, got %v", w.V, g.V)
					}
				case Mul:
					w := iw.(Mul)
					if w.L != g.L {
						t.Fatalf("unexpected locator: want %v, got %v", w.L, g.L)
					}
					if !cmpValue(w.V, g.V) {
						t.Fatalf("unexpected value: want %v, got %v", w.V, g.V)
					}
				case Mod:
					w := iw.(Mod)
					if w.L != g.L {
						t.Fatalf("unexpected locator: want %v, got %v", w.L, g.L)
					}
					if !cmpValue(w.V, g.V) {
						t.Fatalf("unexpected value: want %v, got %v", w.V, g.V)
					}
				case Rcv:
					w := iw.(Rcv)
					if !cmpValue(w.V, g.V) {
						t.Fatalf("unexpected value: want %v, got %v", w.V, g.V)
					}
				case Jgz:
					w := iw.(Jgz)
					if !cmpValue(w.C, g.C) {
						t.Fatalf("unexpected condition: want %v, got %v", w.C, g.C)
					}
					if !cmpValue(w.O, g.O) {
						t.Fatalf("unexpected offset: want %v, got %v", w.O, g.O)
					}
				}
			}
		})
	}
}

