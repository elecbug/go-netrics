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

func (g *GraphParams) ShortestPath(start, end NodeID) Path {
	return &PathParams{algorithm.ShortestPath(g.Graph, start, end)}
}

const INF = graph.INF

const DIRECTED_UNWEIGHTED = graph.DIRECTED_UNWEIGHTED
const DIRECTED_WEIGHTED = graph.DIRECTED_WEIGHTED
const UNDIRECTED_UNWEIGHTED = graph.UNDIRECTED_UNWEIGHTED
const UNDIRECTED_WEIGHTED = graph.UNDIRECTED_WEIGHTED
