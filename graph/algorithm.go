package graph

import (
	"math"
)

func (g Graph) ShortestPath(start, end Identifier) (uint, []Identifier) {
	if g.graphType == DirectedWeighted || g.graphType == UndirectedWeighted {
		return weightedShortestPath(g.ToMatrix(), start, end)
	} else if g.graphType == DirectedUnweighted || g.graphType == UndirectedUnweighted {
		return unweightedShortestPath(g.ToMatrix(), start, end)
	} else {
		return math.MaxUint, nil
	}
}

func weightedShortestPath(m Matrix, start, end Identifier) (uint, []Identifier) {
	const inf uint = math.MaxUint
	n := len(m)

	if int(start) >= n || int(end) >= n {
		return inf, nil
	}

	dist := make([]uint, n)
	prev := make([]int, n)
	visited := make([]bool, n)

	for i := range dist {
		dist[i] = inf
		prev[i] = -1
	}

	dist[start] = 0

	for {
		minDist := inf
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
			if m[u][v] > 0 && !visited[v] {
				alt := dist[u] + m[u][v]
				if alt < dist[v] {
					dist[v] = alt
					prev[v] = u
				}
			}
		}
	}

	path := []Identifier{}

	for at := int(end); at != -1; at = prev[at] {
		path = append(path, Identifier(at))
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	if dist[end] == inf {
		return inf, nil
	}

	return dist[end], path
}

func unweightedShortestPath(m Matrix, start, end Identifier) (uint, []Identifier) {
	const inf uint = math.MaxUint
	n := len(m)

	if int(start) >= n || int(end) >= n {
		return inf, nil
	}

	dist := make([]uint, n)
	prev := make([]int, n)

	for i := range dist {
		dist[i] = inf
		prev[i] = -1
	}

	queue := []int{int(start)}
	dist[start] = 0

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		for v := 0; v < n; v++ {
			if m[u][v] == 0 && dist[v] == inf {
				dist[v] = dist[u] + 1
				prev[v] = u
				queue = append(queue, v)
			}
		}
	}

	path := []Identifier{}

	for at := int(end); at != -1; at = prev[at] {
		path = append(path, Identifier(at))
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	if dist[end] == inf {
		return inf, nil
	}

	return dist[end], path
}
