package vector

import (
	"testing"
)

func TestRange_Each(t *testing.T) {
	tests := map[string]struct {
		in   Range
		want []Vector
	}{
		"0,0 -> 0,0": {
			Range{
				Vector{0,0},
				Vector{0,0},
			},
			[]Vector{
				{0,0},
			},
		},
		"0,0 -> 2,2": {
			Range{
				Vector{0,0},
				Vector{2,2},
			},
			[]Vector{
				{0,0}, {0,1}, {0,2},
				{1,0}, {1,1}, {1,2},
				{2,0}, {2,1}, {2,2},
			},
		},
		"0,0,0 -> 2,2,2": {
			Range{
				Vector{0,0,0},
				Vector{2,2,2},
			},
			[]Vector{
				{0,0,0}, {0,0,1}, {0,0,2}, {0,1,0}, {0,1,1}, {0,1,2}, {0,2,0}, {0,2,1}, {0,2,2},
				{1,0,0}, {1,0,1}, {1,0,2}, {1,1,0}, {1,1,1}, {1,1,2}, {1,2,0}, {1,2,1}, {1,2,2},
				{2,0,0}, {2,0,1}, {2,0,2}, {2,1,0}, {2,1,1}, {2,1,2}, {2,2,0}, {2,2,1}, {2,2,2},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := make([]Vector, 0, len(test.want))
			if err := test.in.Each(func (v Vector) error {
				got = append(got, v)
				return nil
			}); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			cmp := func (a, b []Vector) bool {
				if len(a) != len(b) {
					return false
				}
				for i := range a {
					if !Cmp(a[i], b[i]) {
						return false
					}
				}
				return true
			}
			if !cmp(test.want, got) {
				t.Fatalf("unexpected vectors: want %v, got %v", test.want, got)
			}
		})
	}
}

