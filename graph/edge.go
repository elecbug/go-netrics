package graph

type Edge struct {
	to     Identifier
	weight uint
}

func (n *Node) addEdge(to Identifier, weight uint) {
	n.edges = append(n.edges, &Edge{
		to:     to,
		weight: weight,
	})
}
