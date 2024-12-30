package graph

type Edge struct {
	to     Identifier
	weight uint
}

func newEdge(to Identifier, weight uint) *Edge {
	return &Edge{
		to:     to,
		weight: weight,
	}
}

func (e Edge) To() Identifier {
	return e.to
}

func (e Edge) Weight() uint {
	return e.weight
}
