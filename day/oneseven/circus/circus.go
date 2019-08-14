package circus

import (
	"fmt"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/iterator"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
	"log"
	"strconv"
)

func toBase36(i int64) string {
	return strconv.FormatInt(i, 36)
}

func fromBase36(s string) int64 {
	i, err := strconv.ParseInt(s, 36, 64)
	if err != nil {
		log.Panicf("error decoding base36 to int64: %v", s)
	}
	return i
}

type Tower struct {
	Name      string
	Weight    int
	Subtowers []*Tower
}

func (t *Tower) ID() int64 { return fromBase36(t.Name) }

type Circus map[string]*Tower

func New(descs ...Descriptor) (Circus, error) {
	// Potentially circular data, so do a two-pass initialize
	// First create tower map
	c := make(Circus, len(descs))
	for _, d := range descs {
		c[d.Id] = &Tower{d.Id, d.Weight, nil}
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

func (c Circus) Node(ui int64) graph.Node { return c[toBase36(ui)] }

func (c Circus) Nodes() graph.Nodes {
	n := make([]graph.Node, 0, len(c))
	for _, t := range c {
		n = append(n, t)
	}
	return iterator.NewOrderedNodes(n)
}

func (c Circus) From(ui int64) graph.Nodes {
	ut := c[toBase36(ui)]
	n := make([]graph.Node, len(ut.Subtowers))
	for i, vt := range ut.Subtowers {
		n[i] = vt
	}
	return iterator.NewOrderedNodes(n)
}

func (c Circus) HasEdgeFromTo(ui, vi int64) bool {
	us, vs := toBase36(ui), toBase36(vi)
	ut, vt := c[us], c[vs]
	if ut == nil || vt == nil {
		return false
	}
	for _, st := range ut.Subtowers {
		if st == vt {
			return true
		}
	}
	return false
}

func (c Circus) HasEdgeBetween(xi, yi int64) bool { return c.HasEdgeFromTo(xi, yi) || c.HasEdgeFromTo(yi, xi) }

func (c Circus) Edge(ui, vi int64) graph.Edge {
	if c.HasEdgeFromTo(ui, vi) {
		return simple.Edge{F: c.Node(ui), T: c.Node(vi)}
	}
	return nil
}

func (c Circus) To(vi int64) graph.Nodes {
	vt := c.Node(vi)
	// TODO: hm, this isn't very efficient is it
	m := make(map[*Tower]struct{})
	for _, ut := range c {
		for _, st := range ut.Subtowers {
			if st == vt {
				m[ut] = struct{}{}
				break
			}
		}
	}
	n := make([]graph.Node, 0, len(m))
	for ut := range m {
		n = append(n, ut)
	}
	return iterator.NewOrderedNodes(n)
}

func (c Circus) Base() (*Tower, error) {
	n, err := topo.Sort(c)
	if err != nil {
		return nil, err
	}
	if len(n) == 0 {
		return nil, nil
	}
	return n[0].(*Tower), nil
}
