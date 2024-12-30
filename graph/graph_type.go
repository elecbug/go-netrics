package graph

type GraphType int

const (
	DirectedUnweighted GraphType = iota
	DirectedWeighted
	UndirectedUnweighted
	UndirectedWeighted
)

func (g GraphType) String() string {
	switch g {
	case DirectedUnweighted:
		return "Directed Unweighted Graph"
	case DirectedWeighted:
		return "Directed Weighted Graph"
	case UndirectedUnweighted:
		return "Undirected Unweighted Graph"
	case UndirectedWeighted:
		return "Undirected Weighted Graph"
	default:
		return "Unknown Graph Type"
	}
}
