package orbit

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/traverse"
	"log"
	"math"
	"strconv"
)

func toId(s string) int64 {
	i, err := strconv.ParseInt(s, 36, 64)
	if err != nil {
		log.Panicf("node id parse error: %v", err)
	}
	return i
}

type Body string

func (b Body) ID() int64 { return toId(string(b)) }

type Orbit struct {
	Body   Body
	Orbits Body
}

func NewDirected(orbs ...Orbit) graph.Directed {
	g := simple.NewDirectedGraph()
	for _, o := range orbs {
		// orbited <- orbiter
		// this way around because it simplifies counting indirect orbits
		e := g.NewEdge(o.Body, o.Orbits)
		g.SetEdge(e)
	}
	return g
}

func NewUndirected(orbs ...Orbit) graph.Undirected {
	g := simple.NewUndirectedGraph()
	for _, o := range orbs {
		e := g.NewEdge(o.Body, o.Orbits)
		g.SetEdge(e)
	}
	return g
}

func CountOrbits(g graph.Directed) int {
	count := 0
	// A direct orbit is an edge between two adjacent bodies
	// An indirect orbit is an path between two non-adjacent bodies
	bodies := g.Nodes()
	for bodies.Next() {
		body := bodies.Node()
		// there's no easy way to fetch visited nodes
		// TODO: add this to graph?
		seen := make(map[Body]struct{})
		search := traverse.DepthFirst{
			Visit: func (n graph.Node) { seen[n.(Body)] = struct{}{} },
		}
		search.Walk(g, body, nil)
		count += len(seen) - 1 // will always see self
	}
	return count
}

func AdjDistanceTo(g graph.Undirected, f, t Body) (int, bool) {
	short := path.DijkstraFrom(f, g)
	p, d := short.To(t.ID())
	if math.IsInf(d, 0) {
		return 0, false
	}
	// path includes self
	// two less transfers to account for orbit edges between
	// f - parent and t - parent
	return len(p) - 3, true
}