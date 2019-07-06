package knot

import (
	"container/ring"
	"reflect"
	"strings"
	"testing"
)

func newIntRing(elems ...int) *ring.Ring {
	r := ring.New(len(elems))
	for _, e := range elems {
		r.Value = e
		r = r.Next()
	}
	return r
}

func ringIntSlice(r *ring.Ring) []int {
	a := make([]int, 0, r.Len())
	r.Do(func (v interface{}) {
		a = append(a, v.(int))
	})
	return a
}

func TestReverser_All(t *testing.T) {
	tests := map[string]struct {
		r []int
		want []int
	}{
		"odd len": {[]int{0, 1, 2, 3, 4}, []int{4, 3, 2, 1, 0}},
		"even len": {[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := newIntRing(test.r...)
			rv := NewReverser(r)
			r = rv.All().Ring()
			got := ringIntSlice(r)
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("unexpected values: want %#v, got %#v", test.want, got)
			}
		})
	}
}

func TestReverser_Sector(t *testing.T) {
	tests := map[string]struct {
		sec []int
		offset int
		count int
		want []int
	}{
		"part, not covers head": {[]int{0, 1, 2, 3, 4, 5}, 1, 4, []int{0, 4, 3, 2, 1, 5}},
		"part, covers head": {[]int{0, 1, 2, 3, 4, 5}, -1, 4, []int{1, 0, 5, 3, 4, 2}},
		"all, not covers head": {[]int{0, 1, 2, 3, 4, 5}, 0, 6, []int{5, 4, 3, 2, 1, 0}},
		"all, covers head": {[]int{0, 1, 2, 3, 4, 5}, -1, 6, []int{4, 3, 2, 1, 0, 5}},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := newIntRing(test.sec...)
			rv := NewReverser(r)
			r = rv.Sector(test.offset, test.count).Ring()
			got := ringIntSlice(r)
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("unexpected values: want %#v, got %#v", test.want, got)
			}
		})
	}
}

func TestKnot(t *testing.T) {
	tests := map[string]struct {
		r []int
		str []int
		want []int
	}{
		"example": {
			[]int{0, 1, 2, 3, 4},
			[]int{3, 4, 1, 5},
			[]int{3, 4, 2, 1, 0},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := newIntRing(test.r...)
			str := make(chan int)
			go func() {
				defer close(str)
				for _, size := range test.str {
					str <- size
				}
			}()
			r = Knot(r, str)
			got := ringIntSlice(r)
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("unexpected values: want %#v, got %#v", test.want, got)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	s := "3,4,1,5"
	rd := strings.NewReader(s)
	got, err := Solve(5, rd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := 12
	if want != got {
		t.Fatalf("unexpected value: want %v, got %v", want, got)
	}
}
