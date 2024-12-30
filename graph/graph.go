package graph

import (
	"fmt"
	"math"

	err "github.com/elecbug/go-graphtric/err"
)

type Graph struct {
	nodes     *Nodes
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

func (g *Graph) AddNode(name string) error {
	err := g.nodes.insert(newNode(g.nowID, name))

	if err != nil {
		return err
	}

	g.nowID++

	return nil
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

func (g *Graph) AddWeightEdge(from, to Identifier, weight uint) error {
	if (g.graphType == DirectedUnweighted || g.graphType == UndirectedUnweighted) && weight != 0 {
		return err.InvalidEdge(g.graphType.String(), fmt.Sprintf("weight: %d", weight))
	}

	if g.nodes.find(from) == nil {
		return err.NotExistNode(from.String())
	}
	if g.nodes.find(to) == nil {
		return err.NotExistNode(to.String())
	}

	g.nodes.find(from).addEdge(g.nodes.find(to), weight)

	return nil
}

func (g *Graph) ToMatrix() [][]uint {
	size := len(g.nodes.values)
	matrix := make([][]uint, size)

	for i := range matrix {
		matrix[i] = make([]uint, size)
		for j := range matrix[i] {
			matrix[i][j] = math.MaxUint
		}
	}

	for from_id, from := range g.nodes.values {
		for _, from_edge := range from.edges {
			to_id := from_edge.to.ID()

			matrix[from_id][to_id] = from_edge.weight
		}
	}

	return matrix
}
