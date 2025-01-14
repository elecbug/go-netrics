package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/elecbug/go-netrics/algorithm"
	"github.com/elecbug/go-netrics/graph"
)

func TestEfficiency(t *testing.T) {
	cap := 30
	g := graph.NewGraph(graph.UndirectedUnweighted, cap)

	for i := 0; i < cap; i++ {
		g.AddNode(fmt.Sprintf("%4d", i))
	}

	// t.Logf("%s\n", spew.Sdump(g))

	for i := 0; i < g.NodeCount()*g.NodeCount()/10; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(i)))
		from := graph.Identifier(r.Intn(g.NodeCount()))

		r = rand.New(rand.NewSource(time.Now().UnixNano() + int64(i*i)))
		to := graph.Identifier(r.Intn(g.NodeCount()))

		// t.Logf("%d - %d", from, to)

		g.AddEdge(from, to)
	}

	pu := algorithm.NewParallelUnit(40)

	t.Logf("local eff: %v\n", pu.LocalEfficiency(g))
	t.Logf("global eff: %v\n", pu.GlobalEfficiency(g))

	u := algorithm.NewUnit()

	t.Logf("local eff: %v\n", u.LocalEfficiency(g))
	t.Logf("global eff: %v\n", u.GlobalEfficiency(g))
}
