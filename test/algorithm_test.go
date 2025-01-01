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
	cap := 200
	g := graph.NewGraph(graph.UndirectedUnweighted, cap)

	for i := 0; i < cap; i++ {
		g.AddNode(fmt.Sprintf("%4d", i))
	}

	t.Logf("%s\n", spew.Sdump(g))

	for i := 0; i < g.Size()*g.Size()/100; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(i)))
		from := graph.Identifier(r.Intn(g.Size()))

		r = rand.New(rand.NewSource(time.Now().UnixNano() + int64(i*i)))
		to := graph.Identifier(r.Intn(g.Size()))

		t.Logf("%d - %d", from, to)

		g.AddEdge(from, to)

		// if err != nil {
		// 	t.Logf("%v", err)
		// }
	}

	// t.Logf("\n%s\n", g.ToMatrix().String())

	s := time.Now()

	pm := algorithm.NewParallelMachine(40)
	diameter, nodes := pm.Diameter(g)
	t.Logf("diameter: %d, nodes: %v\n", diameter, nodes)

	duration := time.Since(s)
	t.Logf("Execution time: %s", duration)

	s = time.Now()

	um := algorithm.NewUniMachine()
	diameter, nodes = um.Diameter(g)
	t.Logf("diameter: %d, nodes: %v\n", diameter, nodes)

	duration = time.Since(s)
	t.Logf("Execution time: %s", duration)

	// dist, nodes := pm.ShortestPath(g, graph.Identifier(0), graph.Identifier(1))
	// t.Logf("dist: %d, nodes: %v\n", dist, nodes)
}
