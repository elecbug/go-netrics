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
	for i := 0; i < g.NodeCount(); i++ {
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
	n := g.NodeCount()
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
	for i := 0; i < g.NodeCount(); i++ {
		centrality[graph.Identifier(i)] = 0
	}

	// Define a result type to collect intermediate centrality counts.
	type result struct {
		node  graph.Identifier
		count float64
	}

	resultChan := make(chan result, g.NodeCount())
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
	n := g.NodeCount()
	if n > 2 {
		for node := range centrality {
			centrality[node] /= float64((n - 1) * (n - 2))
		}
	}

	return centrality
}

// DegreeCentrality computes the degree centrality of each node in the graph for a Unit.
// Degree centrality is the number of direct connections a node has to other nodes.
// Parameters:
//   - g: The graph to compute the degree centrality for.
//
// Returns:
//   - A map where the keys are node identifiers and the values are the degree centrality scores.
func (u *Unit) DegreeCentrality(g *graph.Graph) map[graph.Identifier]float64 {
	if !g.Updated() || !u.updated {
		// Recompute shortest paths if the graph or unit has been updated.
		u.computePaths(g)
	}

	centrality := make(map[graph.Identifier]float64)

	// Initialize centrality scores for all nodes to 0.
	for i := 0; i < g.NodeCount(); i++ {
		centrality[graph.Identifier(i)] = 0
	}

	// Calculate the degree for each node by counting direct neighbors.
	matrix := g.ToMatrix()
	for i, row := range matrix {
		for _, value := range row {
			if value != graph.INF {
				centrality[graph.Identifier(i)]++
			}
		}
	}

	// Normalize centrality scores by the maximum possible degree (n-1).
	n := g.NodeCount()
	if n > 1 {
		for node := range centrality {
			centrality[node] /= float64(n - 1)
		}
	}

	return centrality
}

// DegreeCentrality computes the degree centrality of each node in the graph for a ParallelUnit.
// The computation is performed in parallel for better performance on larger graphs.
// Parameters:
//   - g: The graph to compute the degree centrality for.
//
// Returns:
//   - A map where the keys are node identifiers and the values are the degree centrality scores.
func (pu *ParallelUnit) DegreeCentrality(g *graph.Graph) map[graph.Identifier]float64 {
	if !g.Updated() || !pu.updated {
		// Recompute shortest paths if the graph or unit has been updated.
		pu.computePaths(g)
	}

	centrality := make(map[graph.Identifier]float64)

	// Initialize centrality scores for all nodes to 0.
	for i := 0; i < g.NodeCount(); i++ {
		centrality[graph.Identifier(i)] = 0
	}

	matrix := g.ToMatrix()
	var wg sync.WaitGroup
	resultChan := make(chan struct {
		node  graph.Identifier
		count float64
	}, g.NodeCount())

	// Compute degree centrality in parallel.
	for i := 0; i < len(matrix); i++ {
		wg.Add(1)
		go func(nodeIndex int) {
			defer wg.Done()
			count := 0.0
			for _, value := range matrix[nodeIndex] {
				if value != graph.INF {
					count++
				}
			}
			resultChan <- struct {
				node  graph.Identifier
				count float64
			}{node: graph.Identifier(nodeIndex), count: count}
		}(i)
	}

	// Close the result channel after all goroutines complete.
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Aggregate results from the result channel.
	for res := range resultChan {
		centrality[res.node] = res.count
	}

	// Normalize centrality scores by the maximum possible degree (n-1).
	n := g.NodeCount()
	if n > 1 {
		for node := range centrality {
			centrality[node] /= float64(n - 1)
		}
	}

	return centrality
}
