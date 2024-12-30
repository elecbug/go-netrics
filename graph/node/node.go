package node

import (
	"github.com/elecbug/go-graphtric/graph/edge"
	"github.com/elecbug/go-graphtric/graph/gtype"
)

type Node struct {
	identifier gtype.Identifier
	Name       string
	edges      []*edge.Edge
	alive      bool
}

func NewNode(identifier gtype.Identifier, name string) *Node {
	return &Node{
		identifier: identifier,
		Name:       name,
		edges:      make([]*edge.Edge, 0),
		alive:      false,
	}
}

func (n *Node) AddEdge(to gtype.Identifier, weight uint) {
	n.edges = append(n.edges, edge.NewEdge(to, weight))
}

func (n Node) ID() gtype.Identifier {
	return n.identifier
}

func (n Node) Edges() []edge.Edge {
	result := make([]edge.Edge, len(n.edges))

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
}
