package maps

import (
	"github.com/phyrwork/goadvent/collect/set"
	"reflect"
	"testing"
)

func TestChannel_Values(t *testing.T) {
	for name, test := range delegateTests {
		t.Run(name, func(t *testing.T) {
			want := set.New(test.v...)
			m := New(test.p...)
			d := Channel{M: m}
			got := make(set.Set, len(test.v))
			for v := range d.Values() {
				got[v] = struct{}{}
			}
			if !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected values: want %v, got %v", want, got)
			}
		})
	}
}

func TestChannel_Keys(t *testing.T) {
	for name, test := range delegateTests {
		t.Run(name, func(t *testing.T) {
			want := set.New(test.k...)
			m := New(test.p...)
			d := Channel{M: m}
			got := make(set.Set, len(test.k))
			for k := range d.Keys() {
				got[k] = struct{}{}
			}
			if !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected keys: want %v, got %v", want, got)
			}
		})
	}
}

func TestChannel_Pairs(t *testing.T) {
	for name, test := range delegateTests {
		t.Run(name, func(t *testing.T) {
			p := make([]interface{}, len(test.p))
			for i := range test.p {
				p[i] = test.p[i]
			}
			want := set.New(p...)
			m := New(test.p...)
			d := Channel{M: m}
			got := make(set.Set, len(test.p))
			for p := range d.Pairs() {
				got[p] = struct{}{}
			}
			if !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected pairs: want %v, got %v", want, got)
			}
		})
	}
}
