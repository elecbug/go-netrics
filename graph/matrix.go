package graph

import (
	"fmt"
)

type Matrix [][]Distance

func (m Matrix) String() string {
	result := ""

	for _, arr := range [][]Distance(m) {
		for _, a := range arr {
			if a != INF {
				result += fmt.Sprintf("%3d ", a)
			} else {
				result += "INF "
			}
		}

		result += "\n"
	}

	return result
}
