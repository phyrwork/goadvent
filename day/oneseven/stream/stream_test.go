package stream

import (
	"strings"
	"testing"
)

var groups = map[string]struct {
	count int
	score int
	chars int
}{
	"{}":                            {1, 1, 0},
	"{{{}}}":                        {3, 6, 0},
	"{{},{}}":                       {3, 5, 0},
	"{{{},{},{{}}}}":                {6, 16, 0},
	"{<a>,<a>,<a>,<a>}":             {1, 1, 4},
	"{<a!>>}":                       {1, 1, 1},
	"{{<ab>},{<ab>},{<ab>},{<ab>}}": {5, 9, 8},
	"{{<!!>},{<!!>},{<!!>},{<!!>}}": {5, 9, 0},
	"{{<a!>},{<a!>},{<a!>},{<ab>}}": {2, 3, 17},
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

func TestGroup_Chars(t *testing.T) {
	for s, want := range groups {
		t.Run(s, func(t *testing.T) {
			r := strings.NewReader(s)
			g := &Group{}
			if err := parser.Parse(r, g); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			got := g.Chars()
			if got != want.chars {
				t.Fatalf("unexpected value: want %v, got %v", want.score, got)
			}
		})
	}
}