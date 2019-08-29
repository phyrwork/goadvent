package bimap

import (
	"reflect"
	"testing"
)

var setTests = map[string]struct{
	k []interface{}
	v []interface{}
	p []struct{K, V interface{}}
}{
	"empty": {
		[]interface{}{},
		[]interface{}{},
		[]struct{K, V interface{}}{},
	},
	"mirror": {
		[]interface{}{1, 2},
		[]interface{}{2, 1},
		[]struct{K, V interface{}}{{1, 2}, {2, 1}},
	},
	"sequence": {
		[]interface{}{1, 2, 3},
		[]interface{}{4, 5, 6},
		[]struct{K, V interface{}}{{1, 4}, {2, 5},  {3, 6}},
	},
}

func TestSet_Values(t *testing.T) {
	for name, test := range setTests {
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
	for name, test := range setTests {
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
	for name, test := range setTests {
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
				t.Fatalf("unexpected pairs: want %v, got %v", want, got)
			}
		})
	}
}
