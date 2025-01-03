package algorithm

import (
	"math"
	"sort"
	"sync"

	"github.com/elecbug/go-graphtric/graph"
)

func ShortestPath(g *graph.Graph, start, end graph.Identifier) (graph.Distance, []graph.Identifier) {
	if g.Type() == graph.DirectedWeighted || g.Type() == graph.UndirectedWeighted {
		return weightedShortestPath(g.ToMatrix(), start, end)
	} else if g.Type() == graph.DirectedUnweighted || g.Type() == graph.UndirectedUnweighted {
		return unweightedShortestPath(g.ToMatrix(), start, end)
	} else {
		return math.MaxUint, nil
	}
}

func (u *Unit) computePaths(g *graph.Graph) {
	u.shortestPaths = []graph.Path{}
	n := len(g.ToMatrix())

	for start := graph.Identifier(0); start < graph.Identifier(n); start++ {
		for end := graph.Identifier(0); end < graph.Identifier(n); end++ {
			if start == end {
				continue
			}

			distance, nodes := ShortestPath(g, start, end)

			if distance != graph.INF {
				u.shortestPaths = append(u.shortestPaths, *graph.NewPath(distance, nodes))
			}
		}
	}

	sort.Slice(u.shortestPaths, func(i, j int) bool {
		return u.shortestPaths[i].Distance() < u.shortestPaths[j].Distance()
	})

	g.Update()
	u.updated = true
}

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

	for i := uint(0); i < workerCount; i++ {
		go func() {
			defer wg.Done()
			for job := range jobChan {
				distance, nodes := ShortestPath(g, job.start, job.end)

				if distance != graph.INF {
					resultChan <- *graph.NewPath(distance, nodes)
				}
			}
		}()
	}

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

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		pu.shortestPaths = append(pu.shortestPaths, result)
	}

	sort.Slice(pu.shortestPaths, func(i, j int) bool {
		return pu.shortestPaths[i].Distance() < pu.shortestPaths[j].Distance()
	})

	g.Update()
	pu.updated = true
}

func weightedShortestPath(matrix graph.Matrix, start, end graph.Identifier) (graph.Distance, []graph.Identifier) {
	n := len(matrix)

	if int(start) >= n || int(end) >= n {
		return graph.INF, nil
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
		return graph.INF, nil
	}

	return dist[end], path
}

func unweightedShortestPath(matrix graph.Matrix, start, end graph.Identifier) (graph.Distance, []graph.Identifier) {
	n := len(matrix)

	if int(start) >= n || int(end) >= n {
		return graph.INF, nil
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
			if matrix[u][v] == 0 && dist[v] == graph.INF {
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
		return graph.INF, nil
	}

	return dist[end], path
}
