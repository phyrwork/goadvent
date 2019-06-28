package jump

import "testing"

var example1 = Program{0, 3, 0, 1, -3}

func TestExec(t *testing.T) {
	tests := map[string]struct {
		p Program
		ef func (Program) Executable
		c int
	}{
		"example 1": {example1, func (p Program) Executable { return NewJumper(p) }, 5 },
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			exe := test.ef(test.p)
			c := Exec(exe)
			if c != test.c {
				t.Fatalf("unexpected count: want %v, got %v", test.c, c)
			}
		})
	}
}


