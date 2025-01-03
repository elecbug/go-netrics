package algorithm

import (
	"sort"
	"sync"

	"github.com/elecbug/go-graphtric/graph"
)

// ShortestPath computes the shortest path between two nodes in a graph.
// Parameters:
//   - g: The graph to perform the computation on.
//   - start: The starting node identifier.
//   - end: The ending node identifier.
//
// Returns:
//   - A graph.Path containing the shortest path and its total distance.
//   - If no path exists, the returned Path has distance INF and an empty node sequence.
func ShortestPath(g *graph.Graph, start, end graph.Identifier) *graph.Path {
	if g.Type() == graph.DirectedWeighted || g.Type() == graph.UndirectedWeighted {
		return weightedShortestPath(g.ToMatrix(), start, end)
	} else if g.Type() == graph.DirectedUnweighted || g.Type() == graph.UndirectedUnweighted {
		return unweightedShortestPath(g.ToMatrix(), start, end)
	} else {
		return graph.NewPath(graph.INF, []graph.Identifier{})
	}
}

// computePaths calculates all shortest paths between every pair of nodes in the graph for a Unit.
// After computation, the `shortestPaths` field in the Unit is updated and sorted by path distance in ascending order.
// Parameters:
//   - g: The graph to perform the computation on.
func (u *Unit) computePaths(g *graph.Graph) {
	u.shortestPaths = []graph.Path{}
	n := len(g.ToMatrix())

	for start := graph.Identifier(0); start < graph.Identifier(n); start++ {
		for end := graph.Identifier(0); end < graph.Identifier(n); end++ {
			if start == end {
				continue
			}

			path := ShortestPath(g, start, end)

			if path.Distance() != graph.INF {
				u.shortestPaths = append(u.shortestPaths, *path)
			}
		}
	}

	// Sort the paths by their total distance.
	sort.Slice(u.shortestPaths, func(i, j int) bool {
		return u.shortestPaths[i].Distance() < u.shortestPaths[j].Distance()
	})

	g.Update()
	u.updated = true
}

// computePaths calculates all shortest paths in parallel for a ParallelUnit.
// After computation, the `shortestPaths` field in the ParallelUnit is updated and sorted by path distance in ascending order.
// Parameters:
//   - g: The graph to perform the computation on.
func (pu *ParallelUnit) computePaths(g *graph.Graph) {
	pu.shortestPaths = []graph.Path{}

	type to struct {
		start graph.Identifier
		end   graph.Identifier
	}

	n := len(g.ToMatrix())

	jobChan := make(chan to)
	resultChan := make(chan graph.Path)
	workerCount := pu.maxCore

	var wg sync.WaitGroup
	wg.Add(int(workerCount))

	// Start worker goroutines to compute paths in parallel.
	for i := uint(0); i < workerCount; i++ {
		go func() {
			defer wg.Done()
			for job := range jobChan {
				path := ShortestPath(g, job.start, job.end)

				if path.Distance() != graph.INF {
					resultChan <- *path
				}
			}
		}()
	}

	// Generate jobs for every pair of nodes.
	go func() {
		for start := 0; start < n; start++ {
			for end := 0; end < n; end++ {
				if start != end {
					jobChan <- to{graph.Identifier(start), graph.Identifier(end)}
				}
			}
		}
		close(jobChan)
	}()

	// Close the result channel after all workers finish.
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect results from workers.
	for result := range resultChan {
		pu.shortestPaths = append(pu.shortestPaths, result)
	}

	// Sort the paths by their total distance.
	sort.Slice(pu.shortestPaths, func(i, j int) bool {
		return pu.shortestPaths[i].Distance() < pu.shortestPaths[j].Distance()
	})

	g.Update()
	pu.updated = true
}

// weightedShortestPath computes the shortest path between two nodes in a weighted graph.
// Uses Dijkstra's algorithm to calculate the path.
// Parameters:
//   - matrix: The adjacency matrix representation of the graph.
//   - start: The starting node identifier.
//   - end: The ending node identifier.
//
// Returns:
//   - A graph.Path containing the shortest path and its total distance.
func weightedShortestPath(matrix graph.Matrix, start, end graph.Identifier) *graph.Path {
	n := len(matrix)

	if int(start) >= n || int(end) >= n {
		return graph.NewPath(graph.INF, []graph.Identifier{})
	}

	dist := make([]graph.Distance, n)
	prev := make([]int, n)
	visited := make([]bool, n)

	for i := range dist {
		dist[i] = graph.INF
		prev[i] = -1
	}

	dist[start] = 0

	for {
		minDist := graph.INF
		u := -1
		for i := 0; i < n; i++ {
			if !visited[i] && dist[i] < minDist {
				minDist = dist[i]
				u = i
			}
		}

		if u == -1 {
			break
		}

		visited[u] = true

		for v := 0; v < n; v++ {
			if matrix[u][v] > 0 && !visited[v] {
				alt := dist[u] + matrix[u][v]
				if alt < dist[v] {
					dist[v] = alt
					prev[v] = u
				}
			}
		}
	}

	path := []graph.Identifier{}

	for at := int(end); at != -1; at = prev[at] {
		path = append(path, graph.Identifier(at))
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	if dist[end] == graph.INF {
		return graph.NewPath(graph.INF, []graph.Identifier{})
	}

	return graph.NewPath(dist[end], path)
}

// unweightedShortestPath computes the shortest path between two nodes in an unweighted graph.
// Uses BFS to calculate the path.
// Parameters:
//   - matrix: The adjacency matrix representation of the graph.
//   - start: The starting node identifier.
//   - end: The ending node identifier.
//
// Returns:
//   - A graph.Path containing the shortest path and its total distance.
func unweightedShortestPath(matrix graph.Matrix, start, end graph.Identifier) *graph.Path {
	n := len(matrix)

	if int(start) >= n || int(end) >= n {
		return graph.NewPath(graph.INF, []graph.Identifier{})
	}

	dist := make([]graph.Distance, n)
	prev := make([]int, n)

	for i := range dist {
		dist[i] = graph.INF
		prev[i] = -1
	}

	queue := []int{int(start)}
	dist[start] = 0

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		for v := 0; v < n; v++ {
			if matrix[u][v] == 1 && dist[v] == graph.INF {
				dist[v] = dist[u] + 1
				prev[v] = u
				queue = append(queue, v)
			}
		}
	}

	path := []graph.Identifier{}

	for at := int(end); at != -1; at = prev[at] {
		path = append(path, graph.Identifier(at))
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	if dist[end] == graph.INF {
		return graph.NewPath(graph.INF, []graph.Identifier{})
	}

	return graph.NewPath(dist[end], path)
}
