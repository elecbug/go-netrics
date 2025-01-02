package algorithm

import (
	"github.com/elecbug/go-graphtric/graph"
)

func (u *Unit) Diameter(g *graph.Graph) graph.Path {
	if !g.Updated() || !u.updated {
		u.computePaths(g)
	}

	var maxPath graph.Path
	var maxDistance graph.Distance = 0

	for _, path := range u.shortestPaths {
		if path.Distance() > maxDistance {
			maxDistance = path.Distance()
			maxPath = path
		}
	}

	return maxPath
}

func (pu *ParallelUnit) Diameter(g *graph.Graph) graph.Path {
	if !g.Updated() || !pu.updated {
		pu.computePaths(g)
	}

	var maxPath graph.Path
	var maxDistance graph.Distance = 0

	for _, path := range pu.shortestPaths {
		if path.Distance() > maxDistance {
			maxDistance = path.Distance()
			maxPath = path
		}
	}

	return maxPath
}
