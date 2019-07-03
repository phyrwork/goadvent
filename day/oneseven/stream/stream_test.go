package stream

import (
	"strings"
	"testing"
)

var groups = map[string]struct {
	count int
	score int
}{
	"{}":                            {1, 1},
	"{{{}}}":                        {3, 6},
	"{{},{}}":                       {3, 5},
	"{{{},{},{{}}}}":                {6, 16},
	"{<a>,<a>,<a>,<a>}":             {1, 1},
	"{<a!>>}":                       {1, 1},
	"{{<ab>},{<ab>},{<ab>},{<ab>}}": {5, 9},
	"{{<!!>},{<!!>},{<!!>},{<!!>}}": {5, 9},
	"{{<a!>},{<a!>},{<a!>},{<ab>}}": {2, 3},
}

func TestGroup_Count(t *testing.T) {
	for s, want := range groups {
		t.Run(s, func(t *testing.T) {
			r := strings.NewReader(s)
			g := &Group{}
			if err := parser.Parse(r, g); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			got := g.Count()
			if got != want.count {
				t.Fatalf("unexpected value: want %v, got %v", want.count, got)
			}
		})
	}
}

func TestGroup_Score(t *testing.T) {
	for s, want := range groups {
		t.Run(s, func(t *testing.T) {
			r := strings.NewReader(s)
			g := &Group{}
			if err := parser.Parse(r, g); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			got := g.Score(1)
			if got != want.score {
				t.Fatalf("unexpected value: want %v, got %v", want.score, got)
			}
		})
	}
}