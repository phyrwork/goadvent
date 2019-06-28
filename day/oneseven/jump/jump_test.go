package jump

import "testing"

func TestExec(t *testing.T) {
	tests := map[string]struct {
		p Program
		ef func (Program) Executable
		c int
	}{
		"example 1, normal": {Program{0, 3, 0, 1, -3}, func (p Program) Executable { return NewJumper(p, Jump) }, 5 },
		"example 1, strange": {Program{0, 3, 0, 1, -3}, func (p Program) Executable { return NewJumper(p, StrangeJump) }, 10 },
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


