package iterator

import (
	"container/ring"
	"reflect"
	"testing"
)

func TestRingIterator(t *testing.T) {
	tests := map[string]int {
		//"empty ring": 0,
		"one elem": 1,
		"two elems": 2,
		"three or more elems": 3,
	}
	for name, count := range tests {
		t.Run(name, func(t *testing.T) {
			r := ring.New(count)
			want := make([]interface{}, count)
			for i := 0; i < count; i++ {
				r.Value = i
				want[i] = i
				r = r.Next()
			}
			it := NewRingIterator(r)
			got := make([]interface{}, 0)
			for it.Next() {
				got = append(got, it.Value())
			}
			if err := it.Err(); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected elems: want %v, got %v", want, got)
			}
		})
	}
}

