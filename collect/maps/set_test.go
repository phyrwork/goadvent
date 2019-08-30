package maps

import (
	"github.com/phyrwork/goadvent/collect/set"
	"reflect"
	"testing"
)

func TestSet_Values(t *testing.T) {
	for name, test := range delegateTests {
		t.Run(name, func(t *testing.T) {
			want := set.New(test.v...)
			m := New(test.p...)
			d := Set{m}
			got := d.Values()
			if !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected values: want %v, got %v", want, got)
			}
		})
	}
}

func TestSet_Keys(t *testing.T) {
	for name, test := range delegateTests {
		t.Run(name, func(t *testing.T) {
			want := set.New(test.k...)
			m := New(test.p...)
			d := Set{m}
			got := d.Keys()
			if !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected keys: want %v, got %v", want, got)
			}
		})
	}
}

func TestSet_Pairs(t *testing.T) {
	for name, test := range delegateTests {
		t.Run(name, func(t *testing.T) {
			p := make([]interface{}, len(test.p))
			for i := range test.p {
				p[i] = test.p[i]
			}
			want := set.New(p...)
			m := New(test.p...)
			d := Set{m}
			got := d.Pairs()
			if !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected pairs: want %v, got %v", want, got)
			}
		})
	}
}
