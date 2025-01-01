package algorithm

import (
	"sync"

	"github.com/elecbug/go-graphtric/graph"
)

func (um UniMachine) Diameter(g *graph.Graph) *graph.Path {
	n := len(g.ToMatrix())

	var maxDistance graph.Distance = 0
	var longestPath []graph.Identifier

	for start := graph.Identifier(0); start < graph.Identifier(n); start++ {
		for end := graph.Identifier(0); end < graph.Identifier(n); end++ {
			if start == end {
				continue
			}

			distance, nodes := ShortestPath(g, start, end)

			if distance == graph.INF {
				continue
			}

			if distance > maxDistance {
				maxDistance = distance
				longestPath = nodes
			}
		}
	}

	return graph.NewPath(maxDistance, longestPath)
}

func (pm ParallelMachine) Diameter(g *graph.Graph) *graph.Path {
	type to struct {
		start graph.Identifier
		end   graph.Identifier
	}

	n := len(g.ToMatrix())

	jobChan := make(chan to)

	resultChan := make(chan graph.Path)
	workerCount := pm.maxCore

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

	var maxDistance graph.Distance = 0
	var longestPath []graph.Identifier

	for result := range resultChan {
		if result.Distance() > maxDistance {
			maxDistance = result.Distance()
			longestPath = result.Nodes()
		}
	}

	return graph.NewPath(maxDistance, longestPath)
}
