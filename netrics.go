package netrics

import (
	"github.com/elecbug/go-netrics/internal/algorithm"
	"github.com/elecbug/go-netrics/internal/graph"
)

type GraphType = graph.GraphType
type Distance = graph.Distance
type Node = graph.Node
type NodeID = graph.NodeID
type Matrix = graph.Matrix

type Unit = algorithm.Unit
type ParallelUnit = algorithm.ParallelUnit

type GraphParams struct{ *graph.Graph }
type PathParams struct{ *graph.Path }

type Graph interface {
	AddNode(name string) (*Node, error)
	RemoveNode(identifier NodeID) error
	FindNode(identifier NodeID) (*Node, error)
	FindNodesByName(name string) ([]*Node, error)
	AddEdge(from, to NodeID) error
	AddWeightEdge(from, to NodeID, distance Distance) error
	RemoveEdge(from, to NodeID) error
	FindEdge(from, to NodeID) (*Distance, error)
	Matrix() Matrix
	String() string
	NodeCount() int
	EdgeCount() int
	Type() GraphType
	IsUpdated() bool
	ToUnit() *Unit
	ToParallelUnit(core uint) *ParallelUnit
}

type Path interface {
	Distance() Distance
	Nodes() []NodeID
}

var _ Graph = (*GraphParams)(nil)
var _ Path = (*PathParams)(nil)

func NewGraph(graphType GraphType, capacity int) Graph {
	return &GraphParams{graph.NewGraph(graphType, capacity)}
}

func (g *GraphParams) ToUnit() *Unit {
	return algorithm.NewUnit(g.Graph)
}

func (g *GraphParams) ToParallelUnit(core uint) *ParallelUnit {
	return algorithm.NewParallelUnit(g.Graph, core)
}

const INF = Distance(graph.INF)

const (
	DIRECTED_UNWEIGHTED   = GraphType(graph.DIRECTED_UNWEIGHTED)
	DIRECTED_WEIGHTED     = GraphType(graph.DIRECTED_WEIGHTED)
	UNDIRECTED_UNWEIGHTED = GraphType(graph.UNDIRECTED_UNWEIGHTED)
	UNDIRECTED_WEIGHTED   = GraphType(graph.UNDIRECTED_WEIGHTED)
)
