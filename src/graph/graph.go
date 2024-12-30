package graph

type Identifier uint
type GraphType int

const (
	DirectedUnweighted GraphType = iota
	DirectedWeighted
	UndirectedUnweighted
	UndirectedWeighted
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

func (g *Graph) AddNode(name string) {
	g.nodes.Insert(newNode(g.nowID, name))
	g.nowID++
}

func (g *Graph) AddEdge() {
}
