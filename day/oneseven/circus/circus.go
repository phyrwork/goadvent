package circus

import "fmt"

type Tower struct {
	Weight int
	Subtowers []*Tower
}

type Circus map[string]*Tower

func New(descs ...Descriptor) (Circus, error) {
	// Potentially circular data, so do a two-pass initialize
	// First create tower map
	c := make(Circus, len(descs))
	for _, d := range descs {
		c[d.Id] = &Tower{d.Weight, nil}
	}
	// Then add subtowers from map
	for _, d := range descs {
		t := c[d.Id]
		t.Subtowers = make([]*Tower, len(d.Subtowers))
		for i, sid := range d.Subtowers {
			st, ok := c[sid]
			if !ok {
				return c, fmt.Errorf("subtower not described: %v", sid)
			}
			t.Subtowers[i] = st
		}
	}
	return c, nil
}