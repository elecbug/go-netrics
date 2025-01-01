package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/elecbug/go-graphtric/algorithm"
	"github.com/elecbug/go-graphtric/graph"
)

func TestAlgorithm(t *testing.T) {
	g := graph.NewGraph(graph.UndirectedUnweighted, 100)

	for i := 0; i < 100; i++ {
		g.AddNode(fmt.Sprintf("%3d", i))
	}

	t.Logf("%s\n", spew.Sdump(g))

	for i := 0; i < g.Size(); i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(i)))
		from := graph.Identifier(r.Intn(g.Size()))

		r = rand.New(rand.NewSource(time.Now().UnixNano() + int64(i*i)))
		to := graph.Identifier(r.Intn(g.Size()))

		t.Logf("%d - %d", from, to)

		err := g.AddEdge(from, to)

		if err != nil {
			t.Logf("%v", err)
		}
	}

	// t.Logf("\n%s\n", g.ToMatrix().String())

	pm := algorithm.NewParallelMachine(100)

	dist, nodes := pm.ShortestPath(g, graph.Identifier(0), graph.Identifier(1))

	t.Logf("dist: %d, nodes: %v\n", dist, nodes)

	diameter, nodes := pm.Diameter(g)

	t.Logf("diameter: %d, nodes: %v\n", diameter, nodes)
}
