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

		// t.Logf("%d - %d", from, to)

		g.AddEdge(from, to)
	}

	// t.Logf("\n%s\n", g.ToMatrix().String())

	s := time.Now()
	pm := algorithm.NewParallelUnit(40)
	path := pm.Diameter(g)
	t.Logf("diameter: %d, nodes: %v\n", path.Distance(), path.Nodes())
	duration := time.Since(s)
	t.Logf("Execution time: %s", duration)

	s = time.Now()
	um := algorithm.NewUnit()
	path = um.Diameter(g)
	t.Logf("diameter: %d, nodes: %v\n", path.Distance(), path.Nodes())
	duration = time.Since(s)
	t.Logf("Execution time: %s", duration)

	s = time.Now()
	path = pm.Diameter(g)
	t.Logf("diameter: %d, nodes: %v\n", path.Distance(), path.Nodes())
	duration = time.Since(s)
	t.Logf("Execution time: %s", duration)

	for i := 0; i < g.Size()*g.Size(); i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(i)))
		from := graph.Identifier(r.Intn(g.Size()))

		r = rand.New(rand.NewSource(time.Now().UnixNano() + int64(i*i)))
		to := graph.Identifier(r.Intn(g.Size()))

		// t.Logf("%d - %d", from, to)

		g.AddEdge(from, to)
	}

	s = time.Now()
	path = pm.Diameter(g)
	t.Logf("diameter: %d, nodes: %v\n", path.Distance(), path.Nodes())
	duration = time.Since(s)
	t.Logf("Execution time: %s", duration)
}
