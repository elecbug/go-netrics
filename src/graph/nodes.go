package graph

type Nodes struct {
	values map[Identifier]*Node
}

func newNodes(cap int) *Nodes {
	return &Nodes{
		values: make(map[Identifier]*Node, cap),
	}
}

func (ns *Nodes) Insert(node *Node) bool {
	if _, exists := ns.values[node.identifier]; exists {
		return false
	} else {
		ns.values[node.identifier] = node

		return true
	}
}

func (ns *Nodes) Remove(identifier Identifier) bool {
	if _, exists := ns.values[identifier]; exists {
		delete(ns.values, identifier)

		return true
	} else {
		return false
	}
}

func (ns *Nodes) Find(identifier Identifier) *Node {
	return ns.values[identifier]
}
