package graph

// GraphType is an enumeration that defines the type of a graph.
// It specifies whether the graph is directed or undirected and whether it is weighted or unweighted.
type GraphType int

// Enumeration values for GraphType.
// These constants represent different types of graphs:
const (
	DirectedUnweighted   GraphType = iota // A directed graph with unweighted edges.
	DirectedWeighted                      // A directed graph with weighted edges.
	UndirectedUnweighted                  // An undirected graph with unweighted edges.
	UndirectedWeighted                    // An undirected graph with weighted edges.
)

// String converts a GraphType value to its string representation.
// This is useful for displaying the graph type in a human-readable format.
func (g GraphType) String() string {
	switch g {
	case DirectedUnweighted:
		return "Directed Unweighted Graph" // Case for directed unweighted graph.
	case DirectedWeighted:
		return "Directed Weighted Graph" // Case for directed weighted graph.
	case UndirectedUnweighted:
		return "Undirected Unweighted Graph" // Case for undirected unweighted graph.
	case UndirectedWeighted:
		return "Undirected Weighted Graph" // Case for undirected weighted graph.
	default:
		return "Unknown Graph Type" // Default case for unrecognized graph types.
	}
}
