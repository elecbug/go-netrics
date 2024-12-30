package graph

import (
	"errors"

	err_msg "github.com/elecbug/go-graphtric/err/graph"
)

type Nodes struct {
	values  map[Identifier]*Node
	nameMap map[string][]Identifier
}

func newNodes(cap int) *Nodes {
	return &Nodes{
		values:  make(map[Identifier]*Node, cap),
		nameMap: make(map[string][]Identifier, cap),
	}
}

func (ns *Nodes) insert(node *Node) error {
	if _, exists := ns.values[node.identifier]; exists {
		return errors.New(err_msg.AlreadyNode(node.identifier.String()))
	} else {
		ns.values[node.identifier] = node

		if ns.nameMap[node.Name] == nil {
			ns.nameMap[node.Name] = make([]Identifier, 0)
		}

		ns.nameMap[node.Name] = append(ns.nameMap[node.Name], node.identifier)

		return nil
	}
}

func (ns *Nodes) remove(identifier Identifier) error {
	if _, exists := ns.values[identifier]; exists {
		name := ns.values[identifier].Name
		delete(ns.values, identifier)

		for i := 0; i < len(ns.nameMap[name]); i++ {
			if ns.nameMap[name][i] == identifier {
				ns.nameMap[name] = append(ns.nameMap[name][:i], ns.nameMap[name][i+1:]...)
				break
			}
		}

		return nil
	} else {
		return errors.New(err_msg.NotExistNode(identifier.String()))
	}
}

func (ns *Nodes) find(identifier Identifier) *Node {
	return ns.values[identifier]
}

func (ns *Nodes) findAll(name string) []*Node {
	ids := ns.nameMap[name]
	var result = make([]*Node, len(ids))

	for i, id := range ids {
		result[i] = ns.values[id]
	}

	return result
}
