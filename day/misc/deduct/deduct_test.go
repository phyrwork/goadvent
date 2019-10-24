package deduct

import (
	"reflect"
	"testing"
)

var example = []Entry{
	{NewSequence(18211), 0, 3},
	{NewSequence(29790), 1, 1},
	{NewSequence(31307), 2, 1},
	{NewSequence(22294), 0, 1},
	{NewSequence(28334), 0, 3},
}

func TestSequence_C(t *testing.T) {
	tests := map[string]struct {
		A Sequence
		B Sequence
		C int
		D int
	}{
		"18211 <- 31728": {NewSequence(31728), NewSequence(18211), 0, 3},
		"29790 <- 31728": {NewSequence(31728), NewSequence(29790), 1, 1},
		"31307 <- 31728": {NewSequence(31728), NewSequence(31307), 2, 1},
		"22294 <- 31728": {NewSequence(31728), NewSequence(22294), 0, 1},
		"28334 <- 31728": {NewSequence(31728), NewSequence(28334), 0, 3},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			c := test.B.C(test.A)
			if c != test.C {
				t.Fatalf("unexpected C: want %v, got %v", test.C, c)
			}
		})
	}
}

func TestSequence_D(t *testing.T) {
	tests := map[string]struct {
		A Sequence
		B Sequence
		C int
		D int
	}{
		"18211 <- 31728": {NewSequence(31728), NewSequence(18211), 0, 3},
		"29790 <- 31728": {NewSequence(31728), NewSequence(29790), 1, 1},
		"31307 <- 31728": {NewSequence(31728), NewSequence(31307), 2, 1},
		"22294 <- 31728": {NewSequence(31728), NewSequence(22294), 0, 1},
		"28334 <- 31728": {NewSequence(31728), NewSequence(28334), 0, 3},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			d := test.B.D(test.A)
			if d != test.D {
				t.Fatalf("unexpected D: want %v, got %v", test.D, d)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	s := DefaultSequenceSet()
	g := example
	got := Solve(s, g)
	want := SequenceSet{
		NewSequence(31728):{},
		NewSequence(31782):{},
		NewSequence(31820):{},
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("unexpected set: want %v, got %v", want, got)
	}
}