package graph

import (
	"github.com/elecbug/go-netrics/core/internal/graph"
)

type Graph struct {
	*graph.Graph
}

type GraphType struct {
	*graph.GraphType
}

type Distance struct {
	*graph.Distance
}

type Path struct {
	*graph.Path
}

type Node struct {
	*graph.Node
}

type NodeID struct {
	*graph.NodeID
}

type Matrix struct {
	*graph.Matrix
}

var _ graph.IGraph = (*Graph)(nil)
var _ graph.IGraphType = (*GraphType)(nil)
var _ graph.IDistance = (*Distance)(nil)
var _ graph.IPath = (*Path)(nil)
var _ graph.INode = (*Node)(nil)
var _ graph.INodeID = (*NodeID)(nil)
var _ graph.IMatrix = (*Matrix)(nil)

func NewGraph(graphType GraphType, capacity int) *Graph {
	return &Graph{
		graph.NewGraph(*graphType.GraphType, capacity)}
}
