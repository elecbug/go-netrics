package graph

type Edge struct {
	to       Identifier
	distance Distance
}

func newEdge(to Identifier, distance Distance) *Edge {
	return &Edge{
		to:       to,
		distance: distance,
	}
}

func (e Edge) To() Identifier {
	return e.to
}

func (e Edge) Distance() Distance {
	return e.distance
}
