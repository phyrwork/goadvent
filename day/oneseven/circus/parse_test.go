package circus

import (
	"strings"
	"testing"
)

var example =
`pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`

func TestParse(t *testing.T) {
	cmpdesc := func(a, b Descriptor) bool {
		if a.Id != b.Id {
			return false
		}
		if a.Weight != b.Weight {
			return false
		}
		if len(a.Subtowers) != len(b.Subtowers) {
			return false
		}
		// TODO: order isn't really important
		for i := range a.Subtowers {
			if a.Subtowers[i] != b.Subtowers[i] {
				return false
			}
		}
		return true
	}
	tests := map[string]struct {
		in string
		want []Descriptor
	}{
		"none": {
			"",
			[]Descriptor{},
		},
		"single, no subtowers, non terminated": {
			"pbga (66)",
			[]Descriptor{{"pbga", 66, nil}},
		},
		"single, no subtowers, terminated": {
			"pbga (66)\n",
			[]Descriptor{{"pbga", 66, nil}},
		},
		"single, one subtower, non terminated": {
			"fwft (72) -> ktlj",
			[]Descriptor{{"fwft", 72, []string{"ktlj"}}},
		},
		"single, multi subtowers, non terminated": {
			"fwft (72) -> ktlj, cntj, xhth",
			[]Descriptor{{"fwft", 72, []string{"ktlj", "cntj", "xhth"}}},
		},
		"multi, no subtowers, non terminated": {
			"pbga (66)\nxhth (57)\nebii (61)",
			[]Descriptor{
				{"pbga", 66, nil},
				{"xhth", 57, nil},
				{"ebii", 61, nil},
			},
		},
		"multi, no subtowers, terminated": {
			"pbga (66)\nxhth (57)\nebii (61)\n",
			[]Descriptor{
				{"pbga", 66, nil},
				{"xhth", 57, nil},
				{"ebii", 61, nil},
			},
		},
		"multi, multi subtowers, non terminated": {
			"padx (45) -> pbga, havc, qoyq\ntknk (41) -> ugml, padx, fwft",
			[]Descriptor{
				{"padx", 45, []string{"pbga", "havc", "qoyq"}},
				{"tknk", 41, []string{"ugml", "padx", "fwft"}},
			},
		},
		"multi, mixed, non terminated": {
			"padx (45) -> pbga, havc, qoyq\ntknk (41) -> ugml, padx, fwft\njptl (61)",
			[]Descriptor{
				{"padx", 45, []string{"pbga", "havc", "qoyq"}},
				{"tknk", 41, []string{"ugml", "padx", "fwft"}},
				{"jptl", 61, nil},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := Parse(strings.NewReader(test.in))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(got) != len(test.want) {
				t.Fatalf("unexpected desc count: want %v, got %v", len(test.want), len(got))
			}
			for i := range got {
				want, got := test.want[i], got[i]
				if !cmpdesc(want, got) {
					t.Fatalf("unexpected desc #%v: want %v, got %v", i, want, got)
				}
			}
		})
	}
}
