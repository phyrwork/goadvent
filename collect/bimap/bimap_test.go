package bimap

import (
	"reflect"
	"testing"
)

func TestBimap_Set(t *testing.T) {
	tests := map[string]struct {
		init []struct{K, V interface{}}
		k, v interface{}
		want []struct{K, V interface{}}
		err  bool
	}{
		"new": {
			[]struct{K, V interface{}}{},
			1, 2,
			[]struct{K, V interface{}}{{1, 2}},
			false,
		},
		"key exists (same)": {
			[]struct{K, V interface{}}{{1, 2}},
			1, 2,
			[]struct{K, V interface{}}{{1, 2}},
			false,
		},
		"key exists (diff)": {
			[]struct{K, V interface{}}{{1, 2}},
			1, 3,
			nil,
			true,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			want := bimapFromSlice(test.want...)
			got := bimapFromSlice(test.init...)
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
					t.Fatalf("unexpected maps: want %#v, got %#v", want, got)
				}
			}
		})
	}
}

