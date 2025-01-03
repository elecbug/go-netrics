package algorithm

import (
	"github.com/elecbug/go-graphtric/graph"
)

func (u *Unit) Diameter(g *graph.Graph) graph.Path {
	if !g.Updated() || !u.updated {
		u.computePaths(g)
	}

	return u.shortestPaths[len(u.shortestPaths)-1]
}

func (pu *ParallelUnit) Diameter(g *graph.Graph) graph.Path {
	if !g.Updated() || !pu.updated {
		pu.computePaths(g)
	}

	return pu.shortestPaths[len(pu.shortestPaths)-1]
}
