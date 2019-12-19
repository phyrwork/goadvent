package crossed2

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

var vectorRegex = regexp.MustCompile(`(U|D|L|R)(\d+)`)

func ReadVector(s string) (Vector, error) {
	m := vectorRegex.FindStringSubmatch(s)
	if len(m) < 3 {
		return Vector{}, fmt.Errorf("not a point vector string: %v", s)
	}
	d, err := strconv.Atoi(m[2])
	if err != nil {
		return Vector{}, fmt.Errorf("atoi error: %v", err)
	}
	var v Vector
	switch m[1] {
	case "U":
		v = Vector{0, 1}
	case "D":
		v = Vector{0, -1}
	case "L":
		v = Vector{-1, 0}
	case "R":
		v = Vector{1, 0}
	default:
		return Vector{}, fmt.Errorf("unknown direction: %v", m[1])
	}
	return v.Mul(d), nil
}

func Read(r io.Reader) ([][]Vector, error) {
	o := make([][]Vector, 0)
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		s := sc.Text()
		w := strings.Split(s, ",")
		l := make([]Vector, len(w))
		for i, w := range w {
			v, err := ReadVector(w)
			if err != nil {
				return nil, fmt.Errorf("vector read error: %v", err)
			}
			l[i] = v
		}
		o = append(o, l)
	}
	if err := sc.Err(); err != nil {
		return nil, fmt.Errorf("scan error: %v", err)
	}
	return o, nil
}
