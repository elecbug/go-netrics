package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/elecbug/go-graphtric/algorithm"
	"github.com/elecbug/go-graphtric/graph"
)

func TestBetweennessCentrality(t *testing.T) {
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
	result := pu.BetweennessCentrality(g)
	t.Logf("betweenness cen: %v\n", result)

	u := algorithm.NewUnit()
	result = u.BetweennessCentrality(g)
	t.Logf("betweenness cen: %v\n", result)
}
