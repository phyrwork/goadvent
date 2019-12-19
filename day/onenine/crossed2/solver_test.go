package crossed2

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	tests := map[int]struct {
		in string
		want [][]Vector
	}{
		1: {
			strings.Join([]string{
				"R75,D30,R83,U83,L12,D49,R71,U7,L72",
				"U62,R66,U55,R34,D71,R55,D58,R83",
			}, "\n"),
			[][]Vector{
				{{75, 0}, {0, -30}, {83, 0}, {0, 83}, {-12, 0}, {0, -49}, {71, 0}, {0, 7}, {-72, 0}},
				{{0, 62}, {66, 0}, {0, 55}, {34, 0}, {0, -71}, {55, 0}, {0, -58}, {83, 0}},
			},
		},
	}
	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := Read(strings.NewReader(test.in))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("unexpected vectors: want %v, got %v", test.want, got)
			}
		})
	}
}
