package graph

type Node struct {
	identifier Identifier
	Name       string
}

func newNode(identifier Identifier, name string) *Node {
	return &Node{
		identifier: identifier,
		Name:       name,
	}
}

func (n Node) ID() Identifier {
	return n.identifier
}
