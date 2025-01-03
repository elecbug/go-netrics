package algorithm

import (
	"github.com/elecbug/go-graphtric/graph"
)

func (u *Unit) AverageShortestPathLength(g *graph.Graph) float64 {
	if !g.Updated() || !u.updated {
		u.computePaths(g)
	}

	var totalDistance graph.Distance = 0
	var pairCount int

	for _, path := range u.shortestPaths {
		totalDistance += path.Distance()
		pairCount++
	}

	if pairCount == 0 {
		return 0
	}

	return float64(totalDistance) / float64(pairCount)
}

func (pu *ParallelUnit) AverageShortestPathLength(g *graph.Graph) float64 {
	if !g.Updated() || !pu.updated {
		pu.computePaths(g)
	}

	var totalDistance graph.Distance = 0
	var pairCount int

	for _, path := range pu.shortestPaths {
		totalDistance += path.Distance()
		pairCount++
	}

	if pairCount == 0 {
		return 0
	}

	return float64(totalDistance) / float64(pairCount)
}

func (u *Unit) PercentileShortestPathLength(g *graph.Graph, percentile float64) graph.Distance {
	if !g.Updated() || !u.updated {
		u.computePaths(g)
	}

	index := int(percentile * float64(len(u.shortestPaths)))

	if index >= len(u.shortestPaths) {
		index = len(u.shortestPaths) - 1
	} else if index < 0 {
		index = 0
	}

	return u.shortestPaths[index].Distance()
}

func (pu *ParallelUnit) PercentileShortestPathLength(g *graph.Graph, percentile float64) graph.Distance {
	if !g.Updated() || !pu.updated {
		pu.computePaths(g)
	}

	index := int(percentile * float64(len(pu.shortestPaths)))

	if index >= len(pu.shortestPaths) {
		index = len(pu.shortestPaths) - 1
	} else if index < 0 {
		index = 0
	}

	return pu.shortestPaths[index].Distance()
}
