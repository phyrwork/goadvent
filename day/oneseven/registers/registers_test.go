package registers

import (
	"strings"
	"testing"
)

var example =
`b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`

func TestSolveMaxEnd(t *testing.T) {
	ans := SolveMaxEnd(strings.NewReader(example))
	if ans.IsError() {
		t.Fatalf("unexpected error: %v", ans)
	}
	want := "1"
	if got := ans.String(); got != want {
		t.Fatalf("unexpected value: want %v, got %v", want, got)
	}
}

func TestSolveMaxAny(t *testing.T) {
	ans := SolveMaxAny(strings.NewReader(example))
	if ans.IsError() {
		t.Fatalf("unexpected error: %v", ans)
	}
	want := "10"
	if got := ans.String(); got != want {
		t.Fatalf("unexpected value: want %v, got %v", want, got)
	}
}
