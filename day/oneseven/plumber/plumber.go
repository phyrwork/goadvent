package plumber

import (
	"github.com/alecthomas/participle"
	"github.com/alecthomas/participle/lexer"
	"github.com/alecthomas/participle/lexer/ebnf"
	"github.com/phyrwork/goadvent/app"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
	"gonum.org/v1/gonum/graph/traverse"
	"io"
)

var grammar = lexer.Must(ebnf.New(`
	Ident = digit { digit } .
	Punct = "<" | "-" | ">" | "," | "\n" .
	Whitespace = " " .
    digit = "0"…"9" .
`))

type Adj struct {
	From int   `@Ident`
	To   []int `"<" "-" ">" @Ident ("," @Ident)*`
}

type adjlist struct {
	Adjs []Adj `(@@ ("\n" @@)* ("\n")?)?`
}

var parser = participle.MustBuild(&adjlist{}, participle.Lexer(grammar), participle.Elide("Whitespace"))

func Parse(r io.Reader) ([]Adj, error) {
	l := adjlist{}
	return l.Adjs, parser.Parse(r, &l)
}

func NewGraph(adjs ...Adj) *simple.UndirectedGraph {
	g := simple.NewUndirectedGraph()
	for _, adj := range adjs {
		for _, to := range adj.To {
			if to == adj.From {
				// simple doesn't support edges to self (will panic)
				// so the choice is either use multi (which is more hassle to build)
				// or just assert that the node is in the graph without an
				// edge to self, since the edge to self doesn't affect the solution
				n := g.Node(int64(to))
				if n == nil {
					g.AddNode(simple.Node(to))
				}
			} else {
				e := simple.Edge{
					F: simple.Node(adj.From),
					T: simple.Node(to),
				}
				g.SetEdge(e)
			}
		}
	}
	return g
}

func GroupSize(g graph.Undirected, id int) int {
	if g.Node(int64(id)) == nil {
		return 0
	}
	n := 0
	df := traverse.DepthFirst{}
	df.Visit = func (graph.Node) { n++ }
	df.Walk(g, simple.Node(int64(id)), nil)
	return n
}

func SolveSize(r io.Reader) app.Solution {
	adjs, err := Parse(r)
	if err != nil {
		return app.NewError(err)
	}
	g := NewGraph(adjs...)
	n := GroupSize(g, 0)
	return app.Int(n)
}

func SolveCount(r io.Reader) app.Solution {
	adjs, err := Parse(r)
	if err != nil {
		return app.NewError(err)
	}
	g := NewGraph(adjs...)
	n := topo.ConnectedComponents(g)
	return app.Int(len(n))
}