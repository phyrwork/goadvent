package captcha

import (
	"strings"
	"testing"
)

var tests = []struct{
	in string
	sum int
 }{
 	{"1122", 3},
	{"1111", 4},
	{"1234", 0},
	{"91212129", 9},
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
			sum := d.Sum()
			if sum != test.sum {
				t.Fatalf("unexpected sum: want %v, got %v", test.sum, sum)
			}
		})
	}
}
