package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/elecbug/go-netrics/algorithm"
	"github.com/elecbug/go-netrics/graph"
)

func TestAverageShortestPathLength(t *testing.T) {
	cap := 200
	g := graph.NewGraph(graph.UndirectedUnweighted, cap)

	for i := 0; i < cap; i++ {
		g.AddNode(fmt.Sprintf("%4d", i))
	}

	// t.Logf("%s\n", spew.Sdump(g))

	for i := 0; i < g.NodeCount()*g.NodeCount()/100; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(i)))
		from := graph.Identifier(r.Intn(g.NodeCount()))

		r = rand.New(rand.NewSource(time.Now().UnixNano() + int64(i*i)))
		to := graph.Identifier(r.Intn(g.NodeCount()))

		// t.Logf("%d - %d", from, to)

		g.AddEdge(from, to)
	}

	pu := algorithm.NewParallelUnit(g, 40)
	dia := pu.Diameter()
	aspl := pu.AverageShortestPathLength()
	pspl := pu.PercentileShortestPathLength(0.5)
	t.Logf("diameter: %d, ASPL: %f, PSPL: %d\n", dia.Distance(), aspl, pspl)

	u := algorithm.NewUnit(g)
	dia = u.Diameter()
	aspl = u.AverageShortestPathLength()
	t.Logf("diameter: %d, ASPL: %f\n", dia.Distance(), aspl)

	for i := 0.0; i < 1; i += 0.1 {
		t.Logf("%f: %d", i, u.PercentileShortestPathLength(i))
	}
}
