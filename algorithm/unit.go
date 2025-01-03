package algorithm

import (
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
