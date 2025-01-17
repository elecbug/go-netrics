package algorithm

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/elecbug/go-netrics/core/internal/graph"
)

func TestCentrality(t *testing.T) {
	cap := 30
	g := graph.NewGraph(graph.UNDIRECTED_UNWEIGHTED, cap)

	for i := 0; i < cap; i++ {
		g.AddNode(fmt.Sprintf("%4d", i))
	}

	// t.Logf("%s\n", spew.Sdump(g))

	for i := 0; i < g.NodeCount()*g.NodeCount()/10; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(i)))
		from := graph.NodeID(r.Intn(g.NodeCount()))

		r = rand.New(rand.NewSource(time.Now().UnixNano() + int64(i*i)))
		to := graph.NodeID(r.Intn(g.NodeCount()))

		// t.Logf("%d - %d", from, to)

		g.AddEdge(from, to)
	}

	pu := NewParallelUnit(g, 40)
	t.Logf("betweenness cen: %v\n", pu.BetweennessCentrality())
	t.Logf("degree cen: %v\n", pu.DegreeCentrality())
	t.Logf("eigenvector cen: %v\n", pu.EigenvectorCentrality(100, 1e-6))

	u := NewUnit(g)
	t.Logf("betweenness cen: %v\n", u.BetweennessCentrality())
	t.Logf("degree cen: %v\n", u.DegreeCentrality())
	t.Logf("degree cen: %v\n", u.DegreeCentrality())
	t.Logf("eigenvector cen: %v\n", u.EigenvectorCentrality(100, 1e-6))
}
