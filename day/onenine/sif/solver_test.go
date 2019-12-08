package sif

import (
	"reflect"
	"strings"
	"testing"
)

// TODO: it does work
//func TestRead(t *testing.T) {
//	in := "123456789012"
//	want := [][][]rune{
//		{
//			{1,2,3},
//			{4,5,6},
//		},
//		{
//			{7,8,9},
//			{0,1,2},
//		},
//	}
//	r := strings.NewReader(in)
//	got, err := Read(r, 2, 3)
//	if err != nil {
//		t.Fatalf("unexpected error: %v", err)
//	}
//	if !reflect.DeepEqual(want, got) {
//		t.Fatalf("unexpected raw: want %v, got %v", want, got)
//	}
//}

func TestSolve1(t *testing.T) {
	in := "123456789012"
	r := strings.NewReader(in)
	got := solve1(r, 2, 3)
	if got.IsError() {
		t.Fatalf("unexpected error: %v", got.String())
	}
	want := "1"
	if got := got.String(); got != want {
		t.Fatalf("unexpected result: want %v, got %v", want, got)
	}
}

func TestLayer_String(t *testing.T) {
	l := Layer{
		{0, 1},
		{1, 0},
	}
	got := l.String()
	want := "01\n10"
	if got != want {
		t.Fatalf("unexpected string: want %v, got %v", want, got)
	}
}

func TestComposite_Flatten(t *testing.T) {
	i := Composite{
		{{0,2},{2,2}},
		{{1,1},{2,2}},
		{{2,2},{1,2}},
		{{0,0},{0,0}},
	}
	got := i.Flatten()
	want := Layer{
		{0, 1},
		{1, 0},
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("unexpected layer: want %v, got %v", want, got)
	}
}
