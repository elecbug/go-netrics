package algorithm

import (
	"sync"

	"github.com/elecbug/go-graphtric/graph"
)

func (u *Unit) BetweennessCentrality(g *graph.Graph) map[graph.Identifier]float64 {
	if !g.Updated() || !u.updated {
		u.computePaths(g)
	}

	centrality := make(map[graph.Identifier]float64)

	for i := 0; i < g.Size(); i++ {
		centrality[graph.Identifier(i)] = 0
	}

	for _, path := range u.shortestPaths {
		nodes := path.Nodes()

		for _, n := range nodes {
			if n != nodes[0] && n != nodes[len(nodes)-1] {
				centrality[n]++
			}
		}
	}

	n := g.Size()

	if n > 2 {
		for node := range centrality {
			centrality[node] /= float64((n - 1) * (n - 2))
		}
	}

	return centrality
}
func (pu *ParallelUnit) BetweennessCentrality(g *graph.Graph) map[graph.Identifier]float64 {
	if !g.Updated() || !pu.updated {
		pu.computePaths(g)
	}

	centrality := make(map[graph.Identifier]float64)

	for i := 0; i < g.Size(); i++ {
		centrality[graph.Identifier(i)] = 0
	}

	type result struct {
		node  graph.Identifier
		count float64
	}

	resultChan := make(chan result, g.Size())
	var wg sync.WaitGroup

	for _, path := range pu.shortestPaths {
		wg.Add(1)

		go func(path graph.Path) {
			defer wg.Done()
			nodes := path.Nodes()

			for _, n := range nodes {
				if n != nodes[0] && n != nodes[len(nodes)-1] {
					resultChan <- result{node: n, count: 1}
				}
			}
		}(path)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for res := range resultChan {
		centrality[res.node] += res.count
	}

	n := g.Size()

	if n > 2 {
		for node := range centrality {
			centrality[node] /= float64((n - 1) * (n - 2))
		}
	}

	return centrality
}
