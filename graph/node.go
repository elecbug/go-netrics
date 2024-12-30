package graph

import "fmt"

type Identifier uint

func (id Identifier) String() string {
	return fmt.Sprintf("%d", id)
}

type Node struct {
	identifier Identifier
	Name       string
	edges      []*Edge
}

func newNode(identifier Identifier, name string) *Node {
	return &Node{
		identifier: identifier,
		Name:       name,
		edges:      make([]*Edge, 0),
	}
}

func (n Node) ID() Identifier {
	return n.identifier
}
