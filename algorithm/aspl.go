package algorithm

import (
	"sort"

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

	distances := []graph.Distance{}

	for _, path := range u.shortestPaths {
		distances = append(distances, path.Distance())
	}

	if len(distances) == 0 {
		return 0
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i] < distances[j]
	})

	index := int(percentile * float64(len(distances)))
	if index >= len(distances) {
		index = len(distances) - 1
	} else if index < 0 {
		index = 0
	}

	return distances[index]
}

func (pu *ParallelUnit) PercentileShortestPathLength(g *graph.Graph, percentile float64) graph.Distance {
	if !g.Updated() || !pu.updated {
		pu.computePaths(g)
	}

	distances := []graph.Distance{}

	for _, path := range pu.shortestPaths {
		distances = append(distances, path.Distance())
	}

	if len(distances) == 0 {
		return 0
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i] < distances[j]
	})

	index := int(percentile * float64(len(distances)))

	if index >= len(distances) {
		index = len(distances) - 1
	} else if index < 0 {
		index = 0
	}

	return distances[index]
}
