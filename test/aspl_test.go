package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/elecbug/go-graphtric/algorithm"
	"github.com/elecbug/go-graphtric/graph"
)

func TestAverageShortestPathLength(t *testing.T) {
	cap := 200
	g := graph.NewGraph(graph.UndirectedUnweighted, cap)

	for i := 0; i < cap; i++ {
		g.AddNode(fmt.Sprintf("%4d", i))
	}

	// t.Logf("%s\n", spew.Sdump(g))

	for i := 0; i < g.Size()*g.Size()/10; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(i)))
		from := graph.Identifier(r.Intn(g.Size()))

		r = rand.New(rand.NewSource(time.Now().UnixNano() + int64(i*i)))
		to := graph.Identifier(r.Intn(g.Size()))

		// t.Logf("%d - %d", from, to)

		g.AddEdge(from, to)
	}

	pu := algorithm.NewParallelUnit(40)
	dia := pu.Diameter(g)
	aspl := pu.AverageShortestPathLength(g)
	pspl := pu.PercentileShortestPathLength(g, 0.5)
	t.Logf("diameter: %d, ASPL: %f, PSPL: %d\n", dia.Distance(), aspl, pspl)

	u := algorithm.NewUnit()
	dia = u.Diameter(g)
	aspl = u.AverageShortestPathLength(g)
	pspl = u.PercentileShortestPathLength(g, 0.5)
	t.Logf("diameter: %d, ASPL: %f, PSPL: %d\n", dia.Distance(), aspl, pspl)
}
