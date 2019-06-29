package maps

import (
	"reflect"
	"testing"
)

type pair struct { K, V interface{} }

func newBimap(p ...pair) *Bimap {
	v := make(map[interface{}]interface{})
	k := make(map[interface{}]interface{})
	for _, p := range p {
		v[p.K], k[p.V] = p.V, p.K
	}
	return &Bimap{v, k}
}

func TestBimap_Set(t *testing.T) {
	tests := map[string]struct {
		init []pair
		k, v interface{}
		want []pair
		err  bool
	}{
		"new": {
			[]pair{},
			1, 2,
			[]pair{{1, 2}},
			false,
		},
		"key exists (same)": {
			[]pair{{1, 2}},
			1, 2,
			[]pair{{1, 2}},
			false,
		},
		"key exists (diff)": {
			[]pair{{1, 2}},
			1, 3,
			nil,
			true,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			want := newBimap(test.want...)
			got := newBimap(test.init...)
			err := got.Set(test.k, test.v)
			if test.err {
				if err == nil {
					t.Fatalf("expected error")
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if !reflect.DeepEqual(want, got) {
					t.Fatalf("unexpected map: want %#v, got %#v", want, got)
				}
			}
		})
	}
}

