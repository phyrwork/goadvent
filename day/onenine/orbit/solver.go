package orbit

import (
	"bufio"
	"fmt"
	"github.com/phyrwork/goadvent/app"
	"io"
	"regexp"
)

var orbitRegexp = regexp.MustCompile(`((\d|\w)+)\)((\d|\w)+)`)

func ParseOrbit(s string) (Orbit, error) {
	m := orbitRegexp.FindStringSubmatch(s)
	if len(m) < 5 {
		return Orbit{}, fmt.Errorf("not a orbit string: %v", s)
	}
	o := m[1]
	b := m[3]
	return Orbit{Body(b), Body(o)}, nil
}

func ReadOrbits(r io.Reader) ([]Orbit, error) {
	a := make([]Orbit, 0)
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)
	for sc.Scan() {
		s := sc.Text()
		o, err := ParseOrbit(s)
		if err != nil {
			return nil, fmt.Errorf("orbit parse error: %v", err)
		}
		a = append(a, o)
	}
	return a, nil
}

func Solve1(r io.Reader) app.Solution {
	o, err := ReadOrbits(r)
	if err != nil {
		return app.Errorf("read error: %v", err)
	}
	g := NewGraph(o...)
	count := CountOrbits(g)
	return app.Int(count)
}
