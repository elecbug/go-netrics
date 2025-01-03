package algorithm

import (
	"sync"

	"github.com/elecbug/go-graphtric/graph"
)

// BetweennessCentrality computes the betweenness centrality of each node in the graph for a Unit.
// Betweenness centrality measures how often a node appears on the shortest paths between pairs of other nodes.
// Parameters:
//   - g: The graph to compute the betweenness centrality for.
//
// Returns:
//   - A map where the keys are node identifiers and the values are the betweenness centrality scores.
func (u *Unit) BetweennessCentrality(g *graph.Graph) map[graph.Identifier]float64 {
	if !g.Updated() || !u.updated {
		// Recompute shortest paths if the graph or unit has been updated.
		u.computePaths(g)
	}

	centrality := make(map[graph.Identifier]float64)

	// Initialize centrality scores for all nodes to 0.
	for i := 0; i < g.Size(); i++ {
		centrality[graph.Identifier(i)] = 0
	}

	// Count how many times each node appears on the shortest paths.
	for _, path := range u.shortestPaths {
		nodes := path.Nodes()

		for _, n := range nodes {
			// Exclude the source and target nodes of the path.
			if n != nodes[0] && n != nodes[len(nodes)-1] {
				centrality[n]++
			}
		}
	}

	// Normalize the centrality scores.
	n := g.Size()
	if n > 2 {
		for node := range centrality {
			centrality[node] /= float64((n - 1) * (n - 2))
		}
	}

	return centrality
}

// BetweennessCentrality computes the betweenness centrality of each node in the graph for a ParallelUnit.
// The computation is performed in parallel for better performance on larger graphs.
// Parameters:
//   - g: The graph to compute the betweenness centrality for.
//
// Returns:
//   - A map where the keys are node identifiers and the values are the betweenness centrality scores.
func (pu *ParallelUnit) BetweennessCentrality(g *graph.Graph) map[graph.Identifier]float64 {
	if !g.Updated() || !pu.updated {
		// Recompute shortest paths if the graph or unit has been updated.
		pu.computePaths(g)
	}

	centrality := make(map[graph.Identifier]float64)

	// Initialize centrality scores for all nodes to 0.
	for i := 0; i < g.Size(); i++ {
		centrality[graph.Identifier(i)] = 0
	}

	// Define a result type to collect intermediate centrality counts.
	type result struct {
		node  graph.Identifier
		count float64
	}

	resultChan := make(chan result, g.Size())
	var wg sync.WaitGroup

	// Compute centrality scores in parallel.
	for _, path := range pu.shortestPaths {
		wg.Add(1)

		go func(path graph.Path) {
			defer wg.Done()
			nodes := path.Nodes()

			for _, n := range nodes {
				// Exclude the source and target nodes of the path.
				if n != nodes[0] && n != nodes[len(nodes)-1] {
					resultChan <- result{node: n, count: 1}
				}
			}
		}(path)
	}

	// Close the result channel after all goroutines complete.
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Aggregate results from the result channel.
	for res := range resultChan {
		centrality[res.node] += res.count
	}

	// Normalize the centrality scores.
	n := g.Size()
	if n > 2 {
		for node := range centrality {
			centrality[node] /= float64((n - 1) * (n - 2))
		}
	}

	return centrality
}
