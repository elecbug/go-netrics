package graph

import (
	err "github.com/elecbug/go-graphtric/err"
	"github.com/elecbug/go-graphtric/graph/gtype"
	"github.com/elecbug/go-graphtric/graph/node"
)

type Nodes struct {
	nodes   map[gtype.Identifier]*node.Node
	nameMap map[string][]gtype.Identifier
}

func newNodes(cap int) *Nodes {
	return &Nodes{
		nodes:   make(map[gtype.Identifier]*node.Node, cap),
		nameMap: make(map[string][]gtype.Identifier, cap),
	}
}

func (ns *Nodes) insert(node *node.Node) error {
	if _, exists := ns.nodes[node.ID()]; exists {
		return err.AlreadyNode(node.ID().String())
	} else {
		ns.nodes[node.ID()] = node

		if ns.nameMap[node.Name] == nil {
			ns.nameMap[node.Name] = make([]gtype.Identifier, 0)
		}

		ns.nameMap[node.Name] = append(ns.nameMap[node.Name], node.ID())

		return nil
	}
}

func (ns *Nodes) remove(identifier gtype.Identifier) error {
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

func (ns *Nodes) find(identifier gtype.Identifier) *node.Node {
	return ns.nodes[identifier]
}

func (ns *Nodes) findAll(name string) []*node.Node {
	ids := ns.nameMap[name]
	var result = make([]*node.Node, len(ids))

	for i, id := range ids {
		result[i] = ns.nodes[id]
	}

	return result
}
