package explosive

import (
	"bytes"
	"strings"
	"testing"
)

type inflateTest struct {
	in      string
	out     bool
	recurse bool
	want    string
	len     int
}

func TestInflate(t *testing.T) {
	tests := map[string]inflateTest {
		"ADVENT": {
			"ADVENT",
			true,
			false,
			"ADVENT",
			6,
		},
		"A(1x5)BC": {
			"A(1x5)BC",
			true,
			false,
			"ABBBBBC",
			7,
		},
		"(3x3)XYZ": {
			"(3x3)XYZ",
			true,
			false,
			"XYZXYZXYZ",
			9,
		},
		"A(2x2)BCD(2x2)EFG": {
			"A(2x2)BCD(2x2)EFG",
			true,
			false,
			"ABCBCDEFEFG",
			11,
		},
		"(6x1)(1x3)A": {
			"(6x1)(1x3)A",
			true,
			false,
			"(1x3)A",
			6,
		},
		"X(8x2)(3x3)ABCY": {
			"X(8x2)(3x3)ABCY",
			true,
			false,
			"X(3x3)ABC(3x3)ABCY",
			18,
		},
		"X(8x2)(3x3)ABCY (recurse)": {
			"X(8x2)(3x3)ABCY",
			true,
			true,
			"XABCABCABCABCABCABCY",
			20,
		},
		"(27x12)(20x12)(13x14)(7x10)(1x12)A (recurse)": {
			"(27x12)(20x12)(13x14)(7x10)(1x12)A",
			false,
			true,
			"",
			241920,
		},
		"(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN (recurse)": {
			"(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN",
			false,
			true,
			"",
			445,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var b *bytes.Buffer
			if test.out {
				b = &bytes.Buffer{}
			}
			r := strings.NewReader(test.in)
			l, err := Inflate(r, b, test.recurse)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if test.out {
				got := string(b.Bytes())
				if got != test.want {
					t.Fatalf("unexpected string: want %v, got %v", test.want, got)
				}
			}
			if l != test.len {
				t.Fatalf("unexpected len: want %v, got %v", test.len, l)
			}
		})
	}
}

