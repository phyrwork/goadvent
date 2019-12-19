package crossed2

import (
	"github.com/phyrwork/goadvent/collect/bimap"
	"github.com/phyrwork/goadvent/grid2"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
)

type Vector = grid2.Vector

type Point = grid2.Point

type Segment = grid2.Segment

type Segments map[Segment]struct{}

func (m Segments) Contains(s Segment) bool {
	if _, ok := m[s]; ok {
		return true
	}
	if _, ok := m[s.Reverse()]; ok {
		return true
	}
	return false
}

func (m Segments) Set(s Segment) {
	if !m.Contains(s) {
		m[s] = struct{}{}
	}
}

func (m Segments) Clear(s Segment) {
	delete(m, s)
	delete(m, s.Reverse())
}

func (m Segments) Eq(n Segments) bool {
	if len(m) != len(n) {
		return false
	}
	for s := range m {
		if !n.Contains(s) {
			return false
		}
	}
	return true
}

type Wire Segments

func (w Wire) Intersect(x Wire) map[Point]struct{} {
	m := make(map[Point]struct{})
	for s := range w {
		for t := range x {
			if u, ok := s.Intersect(t); ok {
				m[u.Start] = struct{}{}
				m[u.End] = struct{}{}
			}
		}
	}
	return m
}

func NewWire(o Point, v ...Vector) Wire {
	w := make(Wire, len(v))
	cur := o
	for _, v := range v {
		end := cur.Sum(v)
		s := Segment{cur, end}
		Segments(w).Set(s)
		cur = end
	}
	return w
}

type Map struct {
	graph graph.Builder
	keys  bimap.Bimap
}

type Mesh struct {
	m bimap.Bimap
	g graph.Undirected
}

func NewMesh(w ...Wire) *Mesh {
	// find all intersection points
	p := make(map[Point]struct{})
	for x := range w {
		for y := range w {
			if x == y {
				// dont intersect with self
				continue
			}
			for q := range w[x].Intersect(w[y]) {
				p[q] = struct{}{}
			}
		}
	}
	// collate segments and split at intersection points
	s := make(Segments)
	for _, w := range w {
		for t := range w {
			s.Set(t)
		}
	}
	split := func (s Segment, p Point) ([2]Segment, bool) {
		if p == s.Start || p == s.End {
			// split would have no effect
			return [2]Segment{}, false
		}
		if !s.Contains(p) {
			// point doesn't split segment
			return [2]Segment{}, false
		}
		return [2]Segment{
			{s.Start, p},
			{p, s.End},
		}, true
	}
	count := 0
	for {
		if len(s) == count {
			// done splitting
			break
		}
		count = len(s)
		for t := range s {
			for p := range p {
				u, ok := split(t, p)
				if !ok {
					continue
				}
				// replace the old segment with the
				// two new ones
				s.Clear(t)
				s.Set(u[0])
				s.Set(u[1])
			}
		}
	}
	// create graph
	q := make(map[Point]struct{})
	for t := range s {
		q[t.Start] = struct{}{}
		q[t.End] = struct{}{}
	}
	g := simple.NewUndirectedGraph()
	m := bimap.New()
	for r := range q {
		n := g.NewNode()
		if err := m.Set(r, n); err != nil {
			panic(err)
		}
		g.AddNode(n)
	}
	for t := range s {
		u, _ := m.Value(t.Start)
		v, _ := m.Value(t.End)
		g.SetEdge(g.NewEdge(u.(graph.Node), v.(graph.Node)))
	}
	return &Mesh{m, g}
}

func (m *Mesh) Points() map[Point]struct{} {
	it := m.g.Nodes()
	n := 0
	if m := it.Len(); m > 0 {
		n = m
	}
	o := make(map[Point]struct{}, n)
	for it.Next() {
		n := it.Node()
		v, ok := m.m.Key(n)
		if !ok {
			panic("point not found")
		}
		p := v.(Point)
		o[p] = struct{}{}
	}
	return o
}

func (m *Mesh) Segments() map[Segment]struct{} {
	it := m.g.Nodes()
	o := make(Segments)
	for it.Next() {
		uid := it.Node().ID()
		it := m.g.From(uid)
		for it.Next() {
			vid := it.Node().ID()
			e := m.g.Edge(uid, vid)
			if e == nil {
				panic("edge not found")
			}
			p, ok := m.m.Key(e.From())
			if !ok {
				panic("from node not found")
			}
			q, ok := m.m.Key(e.To())
			if !ok {
				panic("to node not found")
			}
			s := Segment{p.(Point), q.(Point)}
			o.Set(s)
		}
	}
	return o
}