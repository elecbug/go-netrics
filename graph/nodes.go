package graph

import (
	err "github.com/elecbug/go-graphtric/err"
)

type Nodes struct {
	nodes   map[Identifier]*Node
	nameMap map[string][]Identifier
}

func newNodes(cap int) *Nodes {
	return &Nodes{
		nodes:   make(map[Identifier]*Node, cap),
		nameMap: make(map[string][]Identifier, cap),
	}
}

func (ns *Nodes) insert(node *Node) error {
	if _, exists := ns.nodes[node.ID()]; exists {
		return err.AlreadyNode(node.ID().String())
	} else {
		ns.nodes[node.ID()] = node

		if ns.nameMap[node.Name] == nil {
			ns.nameMap[node.Name] = make([]Identifier, 0)
		}

		ns.nameMap[node.Name] = append(ns.nameMap[node.Name], node.ID())

		return nil
	}
}

func (ns *Nodes) remove(identifier Identifier) error {
	if _, exists := ns.nodes[identifier]; exists {
		name := ns.nodes[identifier].Name
		delete(ns.nodes, identifier)

		for i := 0; i < len(ns.nameMap[name]); i++ {
			if ns.nameMap[name][i] == identifier {
				ns.nameMap[name] = append(ns.nameMap[name][:i], ns.nameMap[name][i+1:]...)
				break
			}
		}

		return nil
	} else {
		return err.NotExistNode(identifier.String())
	}
}

func (ns *Nodes) find(identifier Identifier) *Node {
	return ns.nodes[identifier]
}

func (ns *Nodes) findAll(name string) []*Node {
	ids := ns.nameMap[name]
	var result = make([]*Node, len(ids))

	for i, id := range ids {
		result[i] = ns.nodes[id]
	}

	return result
}
