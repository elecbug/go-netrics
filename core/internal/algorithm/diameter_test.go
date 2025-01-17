package algorithm

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/elecbug/go-netrics/core/internal/graph"
)

func TestDiameter(t *testing.T) {
	cap := 200
	g := graph.NewGraph(graph.UNDIRECTED_UNWEIGHTED, cap)

	for i := 0; i < cap; i++ {
		g.AddNode(fmt.Sprintf("%4d", i))
	}

	t.Logf("%s\n", spew.Sdump(g))

	for i := 0; i < g.NodeCount()*g.NodeCount()/100; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(i)))
		from := graph.NodeID(r.Intn(g.NodeCount()))

		r = rand.New(rand.NewSource(time.Now().UnixNano() + int64(i*i)))
		to := graph.NodeID(r.Intn(g.NodeCount()))

		// t.Logf("%d - %d", from, to)

		g.AddEdge(from, to)
	}

	// t.Logf("\n%s\n", g.ToMatrix().String())

	s := time.Now()
	pm := NewParallelUnit(g, 40)
	path := pm.Diameter()
	t.Logf("diameter: %d, nodes: %v\n", path.Distance(), path.Nodes())
	duration := time.Since(s)
	t.Logf("execution time: %s", duration)

	s = time.Now()
	um := NewUnit(g)
	path = um.Diameter()
	t.Logf("diameter: %d, nodes: %v\n", path.Distance(), path.Nodes())
	duration = time.Since(s)
	t.Logf("execution time: %s", duration)

	s = time.Now()
	path = pm.Diameter()
	t.Logf("diameter: %d, nodes: %v\n", path.Distance(), path.Nodes())
	duration = time.Since(s)
	t.Logf("execution time: %s", duration)

	for i := 0; i < g.NodeCount()*g.NodeCount(); i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(i)))
		from := graph.NodeID(r.Intn(g.NodeCount()))

		r = rand.New(rand.NewSource(time.Now().UnixNano() + int64(i*i)))
		to := graph.NodeID(r.Intn(g.NodeCount()))

		// t.Logf("%d - %d", from, to)

		g.AddEdge(from, to)
	}

	s = time.Now()
	path = pm.Diameter()
	t.Logf("diameter: %d, nodes: %v\n", path.Distance(), path.Nodes())
	duration = time.Since(s)
	t.Logf("execution time: %s", duration)
}
