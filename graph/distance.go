package graph

import (
	"math"
)

type Distance uint

const INF = Distance(math.MaxUint)

func (w Distance) Int() int {
	return int(w.uint())
}

func (w Distance) uint() uint {
	return uint(w)
}
