package graph

import (
	"fmt"
)

// Identifier represents a unique identifier assigned to nodes in a graph.
// It is defined as an unsigned integer type to ensure non-negative values.
type Identifier uint

// String converts the Identifier to its string representation.
// This is useful for displaying the node's unique identifier in a readable format.
func (id Identifier) String() string {
	// Use fmt.Sprintf to format the Identifier as a decimal string.
	return fmt.Sprintf("%d", id)
}
