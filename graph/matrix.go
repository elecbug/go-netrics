package graph

import (
	"fmt"
)

// Matrix represents the adjacency matrix of a graph.
// Each element in the matrix corresponds to the distance between two nodes.
// If two nodes are not directly connected, the value is set to `INF`.
type Matrix [][]Distance

// String returns a string representation of the Matrix.
// This method formats the matrix for easy readability:
//   - Each row of the matrix is printed on a new line.
//   - Values are separated by spaces, with "INF" used for unreachable nodes.
func (g Graph) String() string {
	result := ""

	matrix := g.ToMatrix()

	// Iterate over each row of the matrix.
	for _, arr := range [][]Distance(matrix) {
		// Iterate over each element in the row.
		for _, a := range arr {
			if a != INF {
				// Print the distance if it is not `INF`.
				result += fmt.Sprintf("%3d ", a)
			} else {
				// Use "INF" to represent unreachable nodes.
				result += "INF "
			}
		}

		// Add a newline at the end of each row.
		result += "\n"
	}

	return result
}
