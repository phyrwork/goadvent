package iterator

import (
	"reflect"
	"strings"
	"testing"
)

func TestRuneIterator(t *testing.T) {
	tests := map[string]struct {
		in []rune
		skip map[rune]struct{}
		want []rune
	}{
		"empty": {
			[]rune{},
			nil,
			[]rune{},
		},
		"no skip": {
			[]rune{0, 1, 2, 3},
			nil,
			[]rune{0, 1, 2, 3},
		},
		"skip whitespace": {
			[]rune{0, 1, '\n', 2, 3, '\n'},
			SkipWhitespace,
			[]rune{0, 1, 2, 3},
		},
		"skip no match": {
			[]rune{0, 1, 2, 3},
			map[rune]struct{}{
				4: {},
			},
			[]rune{0, 1, 2, 3},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			r := strings.NewReader(string(test.in))
			it := NewRuneScanner(r)
			it.Skip = test.skip
			got, err := RuneSlice(it)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("unexpected values: want %v, got %v", test.want, got)
			}
		})
	}
}

