package keypad

import (
	"reflect"
	"strings"
	"testing"
)

func TestDecoder(t *testing.T) {
	tests := []struct {
		name string
		gr   Grammar
		kp   Keypad
		pos  Position
		inst string
		want string
	}{
		{
			"part 1 example",
			DefaultGrammar,
			SquareKeypad,
			Position{1, 1},
			"ULL\nRRDDD\nLURDL\nUUUUD\n",
			"1985",
		},
		{
			"part 2 example",
			DefaultGrammar,
			DiamondKeypad,
			Position{0, 2},
			"ULL\nRRDDD\nLURDL\nUUUUD\n",
			"5DB3",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dc := NewDecoder(test.gr, test.kp, test.pos)
			r := strings.NewReader(test.inst)
			if err := dc.Decode(r); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got := dc.Out(); !reflect.DeepEqual([]rune(test.want), got) {
				t.Fatalf("unexpected out: want %v, got %v", []rune(test.want), got)
			}
		})
	}
}

