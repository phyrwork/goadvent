package twisty

import (
	"bufio"
	"fmt"
	"github.com/phyrwork/goadvent/app"
	"github.com/phyrwork/goadvent/collect/bimap"
	"github.com/phyrwork/goadvent/vector"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/iterator"
	"gonum.org/v1/gonum/graph/path"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/traverse"
	"io"
	"log"
	"math/bits"
	"strconv"
)

type Coord [2]int

const (
	X = 0
	Y = 1
)

var (
	Up = Coord{0, -1}
	Down = Coord{0, 1}
	Left = Coord{-1, 0}
	Right = Coord{1, 0}
)

func (c Coord) Move(d Coord) Coord {
	return Coord{c[X] + d[X], c[Y] + d[Y]}
}

// GeneratorFunc returns true if path is blocked
type GeneratorFunc func (c Coord) bool

func NewGenerate(fav int) GeneratorFunc {
	return func (c Coord) bool {
		x, y := c[X], c[Y]
		// negative co-ordinates are invalid
		if x < 0 {
			return true
		}
		if y < 0 {
			return true
		}
		// generate valid spaces
		s := x*x + 3*x + 2*x*y + y + y*y
		s += fav
		b := bits.OnesCount(uint(s))
		return b % 2 != 0
	}
}

type Graph struct {
	generate GeneratorFunc
	nextId   int64
	id       bimap.Bimap // Coord <-> int64
}

func NewGraph(generate GeneratorFunc) *Graph {
	return &Graph{
		generate: generate,
		nextId:   0,
		id:       bimap.New(),
	}
}

func (g *Graph) CoordNode(c Coord) graph.Node {
	v, ok := g.id.Value(c)
	var cid int64
	if !ok {
		// generate an id
		cid = g.nextId
		g.nextId++
		if err := g.id.Set(c, cid); err != nil {
			log.Panicf("bimap set error: %v", err)
		}
	} else {
		cid = v.(int64)
	}
	return simple.Node(cid)
}

func (g *Graph) Nodes() graph.Nodes {
	o := make([]graph.Node, 0, g.id.Len())
	ids := bimap.Channel{g.id, 20}.Values()
	for v := range ids {
		id := v.(int64)
		o = append(o, simple.Node(id))
	}
	return iterator.NewOrderedNodes(o)
}

func (g *Graph) From(id int64) graph.Nodes {
	k, ok := g.id.Key(id)
	if !ok {
		// not generated
		return nil
	}
	f := k.(Coord)
	if g.generate(f) {
		// is blocked
		return nil
	}
	o := make([]graph.Node, 0, 4)
	for _, d := range []Coord{Up, Down, Left, Right} {
		t := f.Move(d)
		if g.generate(t) {
			// blocked
			continue
		}
		o = append(o, g.CoordNode(t))
	}
	return iterator.NewOrderedNodes(o)
}

func (g *Graph) Edge(uid, vid int64) graph.Edge {
	if _, ok := g.id.Key(uid); !ok {
		return nil
	}
	if _, ok := g.id.Key(vid); !ok {
		return nil
	}
	return simple.Edge{
		F: simple.Node(uid),
		T: simple.Node(vid),
	}
}

func (g *Graph) HasEdgeBetween(xid, yid int64) bool {
	log.Panicf("not implemented")
	return false
}

func (g *Graph) Node(id int64) graph.Node {
	_, ok := g.id.Key(id)
	if !ok {
		return nil
	}
	return simple.Node(id)
}

func (g *Graph) HeuristicCost(x, y graph.Node) float64 {
	var f, t Coord
	xid := x.ID()
	if k, ok := g.id.Key(xid); ok {
		f = k.(Coord)
	} else {
		log.Panicf("node not generated: %v", xid)
	}
	yid := y.ID()
	if k, ok := g.id.Key(yid); ok {
		t = k.(Coord)
	} else {
		log.Panicf("node not generated: %v", yid)
	}
	return float64(vector.Manhattan(f[:], t[:]))
}

func Shortest(g *Graph, f Coord, t Coord) []Coord {
	// generate start and end nodes
	s := g.CoordNode(f)
	e := g.CoordNode(t)
	// search
	short, _ := path.AStar(s, e, g, nil)
	nodes, _ := short.To(e.ID())
	route := make([]Coord, len(nodes))
	for i, n := range nodes {
		nid := n.ID()
		k, ok := g.id.Key(nid)
		if !ok {
			log.Panicf("node not generated: %v", nid)
		}
		route[i] = k.(Coord)
	}
	return route
}

func NewShortestSolver(f Coord, t Coord) app.SolverFunc {
	return func (r io.Reader) (string, error) {
		// TODO: conflate solver setup code
		sc := bufio.NewScanner(r)
		sc.Split(bufio.ScanWords)
		if !sc.Scan() {
			if err := sc.Err(); err != nil {
				return "", fmt.Errorf("scan error: %v", err)
			} else {
				return "", fmt.Errorf("empty scan")
			}
		}
		fav, err := strconv.Atoi(sc.Text())
		if err != nil {
			return "", fmt.Errorf("int decode error: %v", err)
		}
		g := NewGraph(NewGenerate(fav))
		route := Shortest(g, f, t)
		if len(route) == 0 {
			return "", fmt.Errorf("route not found")
		}
		return strconv.Itoa(len(route) - 1), nil
	}
}

func BreadthFirst(g *Graph, f Coord, d int) map[Coord]int {
	// generate start and end nodes
	s := g.CoordNode(f)
	// visit is a hacky way of counting the number of nodes encountered
	// before each depth using the 'until' functionality of DepthFirst
	// TODO: I am not proud of this, but it worked
	seen := make(map[Coord]int)
	seen[f] = 0
	visit := func (n graph.Node, e int) bool {
		nid := n.ID()
		k, ok := g.id.Key(nid)
		if !ok {
			log.Panicf("node not generated: %v", nid)
		}
		c := k.(Coord)
		if f, ok := seen[c]; (!ok || e < f) && e <= d {
			seen[c] = e
		}
		return false // never stop early
	}
	search := traverse.BreadthFirst{}
	search.Walk(g, s, visit)
	return seen
}

func NewSearchSolver(f Coord, d int) app.SolverFunc {
	return func (r io.Reader) (string, error) {
		// TODO: conflate solver setup code
		sc := bufio.NewScanner(r)
		sc.Split(bufio.ScanWords)
		if !sc.Scan() {
			if err := sc.Err(); err != nil {
				return "", fmt.Errorf("scan error: %v", err)
			} else {
				return "", fmt.Errorf("empty scan")
			}
		}
		fav, err := strconv.Atoi(sc.Text())
		if err != nil {
			return "", fmt.Errorf("int decode error: %v", err)
		}
		g := NewGraph(NewGenerate(fav))
		seen := BreadthFirst(g, f, d)
		return strconv.Itoa(len(seen)), nil
	}
}