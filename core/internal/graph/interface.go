package graph

type IGraph interface {
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
}

var _ IGraph = (*Graph)(nil)

type IGraphType interface {
	String() string
}

var _ IGraphType = (*GraphType)(nil)

type IDistance interface {
	ToInt() int
}

var _ IDistance = (*Distance)(nil)

type IPath interface {
	Distance() Distance
	Nodes() []NodeID
}

var _ IPath = (*Path)(nil)

type INode interface {
	ID() NodeID
}

var _ INode = (*Node)(nil)

type INodeID interface {
	String() string
}

var _ INodeID = (*NodeID)(nil)

type IMatrix interface {
}

var _ IMatrix = (*Matrix)(nil)
