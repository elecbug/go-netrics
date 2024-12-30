package edge

import "github.com/elecbug/go-graphtric/graph/gtype"

type Edge struct {
	to     gtype.Identifier
	weight uint
}

func NewEdge(to gtype.Identifier, weight uint) *Edge {
	return &Edge{
		to:     to,
		weight: weight,
	}
}

func (e Edge) To() gtype.Identifier {
	return e.to
}

func (e Edge) Weight() uint {
	return e.weight
}
