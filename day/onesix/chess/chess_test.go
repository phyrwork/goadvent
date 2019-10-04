package chess

import (
	"strings"
	"testing"
)

func TestPasswordGenerator(t *testing.T) {
	sub := NewHashGenerator(HashMd5, "abc", 0)
	gen := NewPasswordGenerator(8, sub, 5)
	want := "18f47a30"
	if !gen.Next() {
		t.Fatal("unexpected end of iterator")
	}
	got := gen.String()
	if want != got {
		t.Fatalf("unexpected password: want %v, got %v", want, got)
	}
}

func TestSolveAppend(t *testing.T) {
	in := "reyedfim\n"
	got, err := SolveAppend(strings.NewReader(in))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "f97c354d"
	if want != got {
		t.Fatalf("unexpected password: wnat %v, got %v", want, got)
	}
}

func TestSolveFiller(t *testing.T) {
	in := "abc\n"
	got, err := SolveFiller(strings.NewReader(in))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "05ace8e3"
	if want != got {
		t.Fatalf("unexpected password: wnat %v, got %v", want, got)
	}
}
