package circus

import (
	"fmt"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/iterator"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
	"gonum.org/v1/gonum/graph/traverse"
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

// Subweight gives the weight of the tower including its subtowers
func (t *Tower) Subweight() int {
	w := 0
	for _, st := range t.Subtowers {
		w += st.Subweight()
	}
	return w + t.Weight
}

func (t *Tower) Balanced() bool {
	w := make(map[int]struct{})
	for _, st := range t.Subtowers {
		w[st.Subweight()] = struct{}{}
	}
	// 0: no subtowers
	// 1: one common subtower weight
	// 2: multiple subtower weights
	return len(w) < 2
}

func (t *Tower) Balance() (string, int, bool, error) {
	w := make(map[int][]*Tower)
	for _, st := range t.Subtowers {
		sw := st.Subweight()
		w[sw] = append(w[sw], st)
	}
	switch len(w) {
	case 0, 1: // Nothing to do
		return "", 0, false, nil
	case 2: // Can balance
		break
	default: // Can't balance
		return "", 0, false, fmt.Errorf(
			"unable to balance tower %v: has more than two (%v) subtower weights", t.Name, len(w))
	}
	var tw, uw int // Target, unique weights
	var tt *Tower  // Target tower (to balance)
	for sw, sts := range w {
		if len(sts) == 1 {
			// Unique weight
			tt, uw = sts[0], sw
		} else {
			tw = sw
		}
	}
	if tt == nil {
		log.Panicf("balance error: nil target tower")
	}
	// Balance
	tt.Weight += tw - uw
	return tt.Name, tt.Weight, true, nil
}

func (t *Tower) Describe() Descriptor {
	d := Descriptor{
		Id: t.Name,
		Weight: t.Weight,
		Subtowers: make([]string, len(t.Subtowers)),
	}
	for i := range t.Subtowers {
		d.Subtowers[i] = t.Subtowers[i].Name
	}
	return d
}

// Circus implements the graph.Directed interface to enable
// a graph top-sort.
//
// Frustratingly, it seems most of the implementation isn't required.
type Circus map[string]*Tower

func NewCircus(descs ...Descriptor) (Circus, error) {
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

func (c Circus) Describe() []Descriptor {
	descs := make([]Descriptor, 0, len(c))
	for _, t := range c {
		descs = append(descs, t.Describe())
	}
	return descs
}

func (c Circus) Clone() (Circus, error) {
	descs := c.Describe()
	return NewCircus(descs...)
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

func (c Circus) Balance() (map[string]int, error) {
	// Starting at the base
	bt, err := c.Base()
	if err != nil {
		return nil, err
	}
	if bt.Balanced() {
		return nil, nil
	}
	// Traverse the graph and find a list of all unbalanced nodes in depth order
	ub := []*Tower{bt}
	df := traverse.BreadthFirst{}
	df.EdgeFilter = func(e graph.Edge) bool { return e.To().(*Tower).Balanced() == false }
	df.Visit = func(_, v graph.Node) { ub = append(ub, v.(*Tower)) }
	df.Walk(c, bt, nil)
	// Work backwards (deepest first) adjusting up to one subtower weight
	// so that the tower is balanced.
	// Keep a log of modified towers to give as the solution.
	mod := make(map[string]int)
	for i := len(ub); i > 0; {
		i--
		t := ub[i]
		// Tower may not need balancing still
		if t.Balanced() {
			continue
		}
		ts, tw, balanced, err := t.Balance()
		if err != nil {
			return nil, err
		}
		if !balanced {
			log.Panicf("didn't balance unbalanced tower")
		}
		mod[ts] = tw
	}
	return mod, nil
}