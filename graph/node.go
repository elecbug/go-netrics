package graph

import (
	"github.com/elecbug/go-netrics/graph/internal/graph_err" // Custom error package
)

// Node represents a node in the graph.
// It contains a unique identifier (`identifier`), a display name (`Name`),
// the edges connected to the node (`edges`), and a flag (`alive`) indicating whether the node is active.
type Node struct {
	identifier Identifier // Unique identifier for the node.
	Name       string     // A human-readable name for the node, which can be duplicated across nodes.
	edges      []*Edge    // A list of edges originating from this node.
}

// newNode creates a new Node instance.
//
// Parameters:
//   - identifier: The unique identifier for the node.
//   - name: The display name for the node.
//
// Returns a pointer to the newly created Node.
func newNode(identifier Identifier, name string) *Node {
	return &Node{
		identifier: identifier,
		Name:       name,
		edges:      make([]*Edge, 0), // Initialize the edges list as empty.
	}
}

// addEdge adds a new edge to the node's list of edges.
//
// Parameters:
//   - to: The identifier of the destination node.
//   - distance: The weight of the edge.
func (n *Node) addEdge(to Identifier, distance Distance) error {
	// Prevent duplicate edges.
	for _, e := range n.edges {
		if e.To() == to {
			return graph_err.AlreadyEdge(n.identifier.String(), to.String())
		}
	}

	n.edges = append(n.edges, newEdge(to, distance))

	return nil
}

// removeEdge removes an edge from the node's list of edges that points to the specified destination node.
//
// Parameters:
//   - to: The identifier of the destination node whose edge needs to be removed.
//
// Returns:
//   - nil if the edge is successfully removed.
//   - An error if the edge does not exist.
//
// Notes:
//   - If the specified edge is found, it is removed, and the node's edge list is updated.
func (n *Node) removeEdge(to Identifier) error {
	for i, e := range n.edges {
		if e.To() == to {
			// Remove the edge by slicing the edge list.
			n.edges = append(n.edges[:i], n.edges[i+1:]...)

			return nil
		}
	}

	// Return an error if the specified edge does not exist.
	return graph_err.NotExistEdge(n.identifier.String(), to.String())
}

// ID returns the unique identifier of the node.
// Useful for accessing or comparing nodes by their identifiers.
func (n Node) ID() Identifier {
	return n.identifier
}

// Edges returns a slice of all edges connected to this node.
// The returned edges are copied from the internal structure to avoid direct modification.
func (n Node) Edges() []Edge {
	result := make([]Edge, len(n.edges)) // Pre-allocate the slice to avoid overhead.

	for _, e := range n.edges {
		result = append(result, *e) // Dereference the pointer to copy the edge.
	}

	return result
}
