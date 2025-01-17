package algorithm

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/elecbug/go-netrics/internal/graph"
)

func TestCoefficient(t *testing.T) {
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

	glo, loc := pu.ClusteringCoefficient()
	t.Logf("clustering coef: %v, %f\n", glo, loc)
	t.Logf("rich club coef: %v\n", pu.RichClubCoefficient(5))

	u := NewUnit(g)

	glo, loc = u.ClusteringCoefficient()
	t.Logf("clustering coef: %v, %f\n", glo, loc)
	t.Logf("rich club coef: %v\n", u.RichClubCoefficient(5))
}
