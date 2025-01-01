package graph

type Node struct {
	identifier Identifier
	Name       string
	edges      []*Edge
	alive      bool
}

func newNode(identifier Identifier, name string) *Node {
	return &Node{
		identifier: identifier,
		Name:       name,
		edges:      make([]*Edge, 0),
		alive:      false,
	}
}

func (n *Node) addEdge(to Identifier, distance Distance) {
	n.edges = append(n.edges, newEdge(to, distance))
}

func (n Node) ID() Identifier {
	return n.identifier
}

func (n Node) Edges() []Edge {
	result := make([]Edge, len(n.edges))

	for _, e := range n.edges {
		result = append(result, *e)
	}

	return result
}

func (n *Node) Up() {
	n.alive = true
}

func (n *Node) Down() {
	n.alive = false
	n.edges = []*Edge{}
}
