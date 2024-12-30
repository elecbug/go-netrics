package graph

import (
	"fmt"
	"math"

	err "github.com/elecbug/go-graphtric/err"
	"github.com/elecbug/go-graphtric/graph/gtype"
	"github.com/elecbug/go-graphtric/graph/node"
)

type Graph struct {
	nodes     *Nodes
	nowID     gtype.Identifier
	graphType gtype.GraphType
}

func NewGraph(graphType gtype.GraphType, capacity int) *Graph {
	return &Graph{
		nodes:     newNodes(capacity),
		nowID:     0,
		graphType: graphType,
	}
}

func (g *Graph) AddNode(name string) error {
	node := node.NewNode(g.nowID, name)
	err := g.nodes.insert(node)

	if err != nil {
		return err
	}

	g.nowID++

	return nil
}

func (g *Graph) RemoveNode(identifier gtype.Identifier) error {
	return g.nodes.remove(identifier)
}

func (g *Graph) FindNode(identifier gtype.Identifier) (*node.Node, error) {
	result := g.nodes.find(identifier)

	if result != nil {
		return result, nil
	} else {
		return nil, err.NotExistNode(identifier.String())
	}
}

func (g *Graph) FindNodesByName(name string) ([]*node.Node, error) {
	result := g.nodes.findAll(name)

	if result != nil {
		return result, nil
	} else {
		return nil, err.NotExistNode(name)
	}
}

func (g *Graph) AddEdge(from, to gtype.Identifier) error {
	return g.AddWeightEdge(from, to, 0)
}

func (g *Graph) AddWeightEdge(from, to gtype.Identifier, weight uint) error {
	if (g.graphType == gtype.DirectedUnweighted || g.graphType == gtype.UndirectedUnweighted) && weight != 0 {
		return err.InvalidEdge(g.graphType.String(), fmt.Sprintf("weight: %d", weight))
	}

	if from == to {
		return err.SelfEdge(from.String())
	}

	if g.nodes.find(from) == nil {
		return err.NotExistNode(from.String())
	}
	if g.nodes.find(to) == nil {
		return err.NotExistNode(to.String())
	}

	for _, e := range g.nodes.find(from).Edges() {
		if e.To() == to {
			return err.AlreadyEdge(from.String(), to.String())
		}
	}

	g.nodes.find(from).AddEdge(to, weight)

	if g.graphType == gtype.UndirectedUnweighted || g.graphType == gtype.UndirectedWeighted {
		g.nodes.find(to).AddEdge(from, weight)
	}

	return nil
}

func (g *Graph) ToMatrix() gtype.Matrix {
	size := g.nowID
	matrix := make([][]uint, size)

	for i := range matrix {
		matrix[i] = make([]uint, size)
		for j := range matrix[i] {
			matrix[i][j] = math.MaxUint
		}
	}

	for from_id, from := range g.nodes.nodes {
		for _, from_edge := range from.Edges() {
			matrix[from_id][from_edge.To()] = from_edge.Weight()
		}
	}

	return matrix
}

func (g Graph) Size() int {
	return len(g.nodes.nodes)
}
