package spinlock

import "testing"

func TestSpinlock_Put(t *testing.T) {
	l := NewSpinlock(3)
	for _, step := range []struct{
		put  int
		next int
	}{
		{1, 0},
		{2, 1},
		{3, 1},
		{4, 3},
		{5, 2},
		{6, 1},
		{7, 2},
		{8, 6},
		{9, 5},
	}{
		l.Put(step.put)
		next := l.Peek(1)
		if next != step.next {
			t.Fatalf("unexpected value: want %v, got %v", step.next, next)
		}
	}
}

func TestSpinlock_Stream(t *testing.T) {
	c := NewSequence(1, 2017)
	l := NewSpinlock(3)
	l.Stream(c)
	want := 638
	got := l.Peek(1)
	if got != want {
		t.Fatalf("unexpected value: want %v, got %v", want, got)
	}
}

