package bimap

import (
	"reflect"
	"testing"
)

func TestSet_Values(t *testing.T) {
	for name, test := range delegateTests {
		t.Run(name, func(t *testing.T) {
			want := setFromSlice(test.v...)
			m := bimapFromSlice(test.p...)
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
			want := setFromSlice(test.k...)
			m := bimapFromSlice(test.p...)
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
			want := setFromSlice(p...)
			m := bimapFromSlice(test.p...)
			d := Set{m}
			got := d.Pairs()
			if !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected delegateTests: want %v, got %v", want, got)
			}
		})
	}
}
