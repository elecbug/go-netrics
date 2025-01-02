package algorithm

import (
	"sync"

	"github.com/elecbug/go-graphtric/graph"
)

type Unit struct {
	shortestPaths []graph.Path
	updated       bool
}

type ParallelUnit struct {
	Unit
	maxCore uint
}

func NewUnit() *Unit {
	return &Unit{
		shortestPaths: make([]graph.Path, 0),
		updated:       false,
	}
}

func NewParallelUnit(core uint) *ParallelUnit {
	return &ParallelUnit{
		Unit:    *NewUnit(),
		maxCore: core,
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

	g.ComputePaths()
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

	g.ComputePaths()
	pu.updated = true
}
