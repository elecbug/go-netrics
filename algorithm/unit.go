package algorithm

import (
	"github.com/elecbug/go-graphtric/graph"
)

// Unit represents a computation unit for graph algorithms.
// It is used to store and compute shortest paths within a graph.
//
// Fields:
//   - shortestPaths: A slice of all shortest paths in the graph, sorted by their distance in ascending order.
//   - updated: A boolean indicating whether the paths are up-to-date or if the graph has been modified.
type Unit struct {
	shortestPaths []graph.Path // Stores the shortest paths for the graph, sorted by distance in ascending order.
	updated       bool         // Indicates whether the data needs to be recalculated.
}

// ParallelUnit is an extension of Unit for parallel computation.
// It supports running graph algorithms using multiple cores.
//
// Fields:
//   - Unit: Embeds the base Unit structure for algorithm computations.
//   - maxCore: The maximum number of cores to be used for parallel processing.
type ParallelUnit struct {
	Unit         // Embeds the Unit structure for shared functionality.
	maxCore uint // Maximum number of CPU cores to use for parallel computation.
}

// NewUnit creates and initializes a new Unit instance.
// Returns a pointer to the newly created Unit.
func NewUnit() *Unit {
	return &Unit{
		shortestPaths: make([]graph.Path, 0), // Initialize with an empty slice of paths.
		updated:       false,                 // Initially set to false, indicating no updates yet.
	}
}

// NewParallelUnit creates and initializes a new ParallelUnit instance.
//
// Parameters:
//   - core: The maximum number of cores to use for parallel computations.
//
// Returns a pointer to the newly created ParallelUnit.
func NewParallelUnit(core uint) *ParallelUnit {
	return &ParallelUnit{
		Unit:    *NewUnit(), // Initialize the embedded Unit structure.
		maxCore: core,       // Set the maximum number of cores for parallel processing.
	}
}
