package orbit

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/traverse"
	"log"
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

func NewGraph(orbs ...Orbit) graph.Directed {
	g := simple.NewDirectedGraph()
	for _, o := range orbs {
		// orbited <- orbiter
		// this way around because it simplifies counting indirect orbits
		e := g.NewEdge(o.Body, o.Orbits)
		g.SetEdge(e)
	}
	return g
}

//func MapDepth(f Body, g graph.Directed) map[Body]int {
//	depth := make(map[Body]int)
//	// visit is a hacky way of collecting node depth by using
//	// until as a visit func
//	visit := func (n graph.Node, d int) bool {
//		b := n.(Body)
//		if _, ok := depth[b]; !ok {
//			depth[b] = d
//		}
//		return false
//	}
//	search := traverse.BreadthFirst{}
//	search.Walk(g, f, visit)
//	return depth
//}

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