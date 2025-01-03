package graph

// Edge represents a connection (edge) between two nodes in a graph.
// It contains information about the destination node (`to`) and the weight of the edge (`distance`).
type Edge struct {
	to       Identifier // The destination node's unique identifier.
	distance Distance   // The weight or cost of traveling along this edge.
}

// newEdge creates a new Edge instance.
// Parameters:
//   - to: The destination node's identifier.
//   - distance: The weight of the edge.
// Returns a pointer to the newly created Edge.
func newEdge(to Identifier, distance Distance) *Edge {
	return &Edge{
		to:       to,
		distance: distance,
	}
}

// To returns the identifier of the destination node for this edge.
// This is useful for accessing the endpoint of the edge.
func (e Edge) To() Identifier {
	return e.to
}

// Distance returns the weight of the edge.
// This represents the cost or distance associated with traveling this edge.
func (e Edge) Distance() Distance {
	return e.distance
}
