package graph

import (
	"fmt"
	"math"

	err "github.com/elecbug/go-graphtric/err"
)

type Graph struct {
	nodes     *graphNodes
	nowID     Identifier
	graphType GraphType
}

func NewGraph(graphType GraphType, capacity int) *Graph {
	return &Graph{
		nodes:     newNodes(capacity),
		nowID:     0,
		graphType: graphType,
	}
}

func (g *Graph) AddNode(name string) (*Node, error) {
	node := newNode(g.nowID, name)
	err := g.nodes.insert(node)

	if err != nil {
		return nil, err
	}

	g.nowID++

	return node, nil
}

func (g *Graph) RemoveNode(identifier Identifier) error {
	return g.nodes.remove(identifier)
}

func (g *Graph) FindNode(identifier Identifier) (*Node, error) {
	result := g.nodes.find(identifier)

	if result != nil {
		return result, nil
	} else {
		return nil, err.NotExistNode(identifier.String())
	}
}

func (g *Graph) FindNodesByName(name string) ([]*Node, error) {
	result := g.nodes.findAll(name)

	if result != nil {
		return result, nil
	} else {
		return nil, err.NotExistNode(name)
	}
}

func (g *Graph) AddEdge(from, to Identifier) error {
	return g.AddWeightEdge(from, to, 0)
}

func (g *Graph) AddWeightEdge(from, to Identifier, distance Distance) error {
	if (g.graphType == DirectedUnweighted || g.graphType == UndirectedUnweighted) && distance != 0 {
		return err.InvalidEdge(g.graphType.String(), fmt.Sprintf("weight: %d", distance))
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

	g.nodes.find(from).addEdge(to, distance)

	if g.graphType == UndirectedUnweighted || g.graphType == UndirectedWeighted {
		g.nodes.find(to).addEdge(from, distance)
	}

	return nil
}

func (g *Graph) ToMatrix() Matrix {
	size := g.nowID
	matrix := make([][]Distance, size)

	for i := range matrix {
		matrix[i] = make([]Distance, size)
		for j := range matrix[i] {
			matrix[i][j] = math.MaxUint
		}
	}

	for from_id, from := range g.nodes.nodes {
		for _, from_edge := range from.Edges() {
			matrix[from_id][from_edge.To()] = from_edge.Distance()
		}
	}

	return matrix
}

func (g Graph) Size() int {
	return len(g.nodes.nodes)
}

func (g Graph) Type() GraphType {
	return g.graphType
}
