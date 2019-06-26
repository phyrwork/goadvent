package captcha

import (
	"strings"
	"testing"
)

var tests = []struct{
	in string
	cmp CmpFunc
	sum int
 }{
 	{"1122", Next, 3},
	{"1111", Next, 4},
	{"1234", Next, 0},
	{"91212129", Next, 9},
	{"1212", Half, 6},
	{"1221", Half, 0},
	{"123425", Half, 4},
	{"123123", Half, 12},
	{"12131415", Half, 4},
}

func TestNewDigits(t *testing.T) {
	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			rd := strings.NewReader(test.in)
			d, err := NewDigits(rd)
			if err != nil {
				t.Fatalf("error creating object under test: %v", err)
			}
			out := d.String()
			if out != test.in {
				t.Fatalf("unexpected digits string: want '%v', got '%v'", test.in, out)
			}
		})
	}
}

func TestDigits_Sum(t *testing.T) {
	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			rd := strings.NewReader(test.in)
			d, err := NewDigits(rd)
			if err != nil {
				t.Fatalf("error creating object under test: %v", err)
			}
			sum := d.Sum(test.cmp)
			if sum != test.sum {
				t.Fatalf("unexpected sum: want %v, got %v", test.sum, sum)
			}
		})
	}
}
