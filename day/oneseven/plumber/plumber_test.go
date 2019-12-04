package plumber

import (
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	tests := map[string]struct {
		in   string
		want []Adj
	}{
		"single, one adj":{
			"1 <-> 430",
			[]Adj{
				{1, []int{430}},
			},
		},
		"single, multi adj":{
			"3 <-> 303, 363, 635",
			[]Adj{
				{3, []int{303, 363, 635}},
			},
		},
		"multi, mixed adj":{
			"0 <-> 1352, 1864\n1 <-> 430",
			[]Adj{
				{0, []int{1352, 1864}},
				{1, []int{430}},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := Parse(strings.NewReader(test.in))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("unexpected adjs: want %v, got %v", test.want, got)
			}
		})
	}
}

var example =
`0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5`

func TestSolveSize(t *testing.T) {
	ans := SolveSize(strings.NewReader(example))
	if ans.IsError() {
		t.Fatalf("unexpected error: %v", ans)
	}
	want := "6"
	if got := ans.String(); want != got {
		t.Fatalf("unexpected size: want %v, got %v", want, got)
	}
}

func TestSolveCount(t *testing.T) {
	ans := SolveCount(strings.NewReader(example))
	if ans.IsError() {
		t.Fatalf("unexpected error: %v", ans)
	}
	want := "2"
	if got := ans.String(); want != got {
		t.Fatalf("unexpected size: want %v, got %v", want, got)
	}
}


