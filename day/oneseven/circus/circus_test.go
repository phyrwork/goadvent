package circus

import "testing"

var example = []Descriptor{
	{"pbga", 66, nil},
	{"xhth", 57, nil},
	{"ebii", 61, nil},
	{"havc", 66, nil},
	{"ktlj", 57, nil},
	{"fwft", 72, []string{"ktlj", "cntj", "xhth"}},
	{"qoyq", 66, nil},
	{"padx", 45, []string{"pbga", "havc", "qoyq"}},
	{"tknk", 41, []string{"ugml", "padx", "fwft"}},
	{"jptl", 61, nil},
	{"ugml", 68, []string{"gyxo", "ebii", "jptl"}},
	{"gyxo", 61, nil},
	{"cntj", 57, nil},
}

func TestCircus_Base(t *testing.T) {
	c, err := NewCircus(example...)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "tknk"
	bt, err := c.Base()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if bt.Name != want {
		t.Fatalf("unexpected base: want %v, got %v", want, bt.Name)
	}
}

func TestCircus_Balance(t *testing.T) {
	c, err := NewCircus(example...)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := map[string]int{"ugml": 60}
	got, err := c.Balance()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != len(want) {
		t.Fatalf("unexpected balance count: want %v, got %v", len(want), len(got))
	}
	for ts, twg := range got {
		tww, ok := want[ts]
		if !ok {
			t.Fatalf("unexpected tower balanced: %v", ts)
		}
		if twg != tww {
			t.Fatalf("unexpected weight, tower %v: want %v, got %v", ts, tww, twg)
		}
	}
}

