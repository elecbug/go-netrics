package graph

type Edge struct {
	to     *Node
	weight uint
}

func (n *Node) addEdge(to *Node, weight uint) {
	n.edges = append(n.edges, &Edge{
		to:     to,
		weight: weight,
	})
}
