package graph

type Path struct {
	distance Distance
	nodes    []Identifier
}

func NewPath(distance Distance, nodes []Identifier) *Path {
	return &Path{
		distance: distance,
		nodes:    nodes,
	}
}

func (p Path) Distance() Distance {
	return p.distance
}

func (p Path) Nodes() []Identifier {
	return p.nodes
}
