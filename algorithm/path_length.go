package algorithm

import (
	"github.com/elecbug/go-graphtric/graph"
)

// AverageShortestPathLength computes the average shortest path length in the graph.
//
// Parameters:
//   - g: The graph to compute the metric for.
//
// Returns:
//   - The average shortest path length as a float64.
//
// Notes:
//   - If no shortest paths are found, the function returns 0.
func (u *Unit) AverageShortestPathLength(g *graph.Graph) float64 {
	if !g.Updated() || !u.updated {
		// Recompute shortest paths if the graph or unit has been updated.
		u.computePaths(g)
	}

	var totalDistance graph.Distance = 0
	var pairCount int

	// Sum up distances for all shortest paths.
	for _, path := range u.shortestPaths {
		totalDistance += path.Distance()
		pairCount++
	}

	if pairCount == 0 {
		return 0 // Avoid division by zero if no paths exist.
	}

	return float64(totalDistance) / float64(pairCount)
}

// ParallelUnit version of AverageShortestPathLength.
// Computes the average shortest path length using parallel computations.
func (pu *ParallelUnit) AverageShortestPathLength(g *graph.Graph) float64 {
	if !g.Updated() || !pu.updated {
		// Recompute shortest paths if the graph or unit has been updated.
		pu.computePaths(g)
	}

	var totalDistance graph.Distance = 0
	var pairCount int

	// Sum up distances for all shortest paths.
	for _, path := range pu.shortestPaths {
		totalDistance += path.Distance()
		pairCount++
	}

	if pairCount == 0 {
		return 0 // Avoid division by zero if no paths exist.
	}

	return float64(totalDistance) / float64(pairCount)
}

// PercentileShortestPathLength returns the shortest path length at the specified percentile.
//
// Parameters:
//   - g: The graph to compute the metric for.
//   - percentile: A float64 between 0 and 1 indicating the desired percentile.
//
// Returns:
//   - The shortest path length corresponding to the given percentile.
//
// Notes:
//   - The percentile is calculated based on the sorted list of shortest paths.
//   - If the percentile is out of range, it is clamped to valid indices.
func (u *Unit) PercentileShortestPathLength(g *graph.Graph, percentile float64) graph.Distance {
	if !g.Updated() || !u.updated {
		// Recompute shortest paths if the graph or unit has been updated.
		u.computePaths(g)
	}

	// Calculate the index for the desired percentile.
	index := int(percentile * float64(len(u.shortestPaths)))

	// Clamp the index to the valid range.
	if index >= len(u.shortestPaths) {
		index = len(u.shortestPaths) - 1
	} else if index < 0 {
		index = 0
	}

	return u.shortestPaths[index].Distance()
}

// ParallelUnit version of PercentileShortestPathLength.
// Computes the percentile shortest path length using parallel computations.
func (pu *ParallelUnit) PercentileShortestPathLength(g *graph.Graph, percentile float64) graph.Distance {
	if !g.Updated() || !pu.updated {
		// Recompute shortest paths if the graph or unit has been updated.
		pu.computePaths(g)
	}

	// Calculate the index for the desired percentile.
	index := int(percentile * float64(len(pu.shortestPaths)))

	// Clamp the index to the valid range.
	if index >= len(pu.shortestPaths) {
		index = len(pu.shortestPaths) - 1
	} else if index < 0 {
		index = 0
	}

	return pu.shortestPaths[index].Distance()
}
