package knot

import (
	"container/ring"
	"golang.org/x/sync/errgroup"
	"reflect"
	"strings"
	"testing"
)

func newByteRing(elems ...byte) *ring.Ring {
	r := ring.New(len(elems))
	for _, e := range elems {
		r.Value = e
		r = r.Next()
	}
	return r
}

func ringByteSlice(r *ring.Ring) []byte {
	a := make([]byte, 0, r.Len())
	r.Do(func (v interface{}) {
		a = append(a, v.(byte))
	})
	return a
}

func newByteStream(e ...byte) StreamFunc {
	return func(out chan <-byte) error {
		go func () {
			defer close(out)
			for _, e := range e {
				out <- e
			}
		}()
		return nil
	}
}

func TestReverser_All(t *testing.T) {
	tests := map[string]struct {
		r []byte
		want []byte
	}{
		"odd len": {[]byte{0, 1, 2, 3, 4}, []byte{4, 3, 2, 1, 0}},
		"even len": {[]byte{1, 2, 3, 4}, []byte{4, 3, 2, 1}},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := newByteRing(test.r...)
			rv := NewReverser(r)
			r = rv.All().Ring()
			got := ringByteSlice(r)
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("unexpected values: want %#v, got %#v", test.want, got)
			}
		})
	}
}

func TestReverser_Sector(t *testing.T) {
	tests := map[string]struct {
		sec []byte
		offset int
		count int
		want []byte
	}{
		"part, not covers head": {[]byte{0, 1, 2, 3, 4, 5}, 1, 4, []byte{0, 4, 3, 2, 1, 5}},
		"part, covers head": {[]byte{0, 1, 2, 3, 4, 5}, -1, 4, []byte{1, 0, 5, 3, 4, 2}},
		"all, not covers head": {[]byte{0, 1, 2, 3, 4, 5}, 0, 6, []byte{5, 4, 3, 2, 1, 0}},
		"all, covers head": {[]byte{0, 1, 2, 3, 4, 5}, -1, 6, []byte{4, 3, 2, 1, 0, 5}},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := newByteRing(test.sec...)
			rv := NewReverser(r)
			r = rv.Sector(test.offset, test.count).Ring()
			got := ringByteSlice(r)
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("unexpected values: want %#v, got %#v", test.want, got)
			}
		})
	}
}

func TestSparseHash(t *testing.T) {
	tests := map[string]struct {
		n    int
		strf StreamFunc
		want []byte
	}{
		"example 1": {
			5,
			newByteStream(3, 4, 1, 5),
			[]byte{3, 4, 2, 1, 0},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			str := make(chan byte)
			g := errgroup.Group{}
			g.Go(func () error {
				return test.strf(str)
			})
			got := SparseHash(test.n, str)
			if err := g.Wait(); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("unexpected values: want %#v, got %#v", test.want, got)
			}
		})
	}
}

func TestKnotHash(t *testing.T) {
	tests := map[string]string {
		"": "a2582a3a0e66e6e86e3812dcb672a272",
		"AoC 2017": "33efeb34ea91902bb2f59c9920caa6cd",
		"1,2,3": "3efbe78a8d82f29979031a4aa0b16a9d",
		"1,2,4": "63960835bcdc130f0b66d7ff4f6a5a8e",
	}
	for key, want := range tests {
		t.Run(key, func(t *testing.T) {
			ans := KnotHash(strings.NewReader(key))
			if ans.IsError() {
				t.Errorf("unexpected error: %v", ans)
			}
			if got := ans.String(); got != want {
				t.Errorf("unexpected value: want %v, got %v", want, got)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	s := "3,4,1,5"
	rd := strings.NewReader(s)
	got, err := solveSparse(5, NewCommaStream(rd))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := 12
	if want != got {
		t.Fatalf("unexpected value: want %v, got %v", want, got)
	}
}
