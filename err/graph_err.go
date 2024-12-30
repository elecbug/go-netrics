package graph_err

import (
	"fmt"
)

func InvalidEdge(graphKey, edgeKey string) error {
	return fmt.Errorf("invalid edge for graph type: [(%s) does not fit %s]", edgeKey, graphKey)
}

func AlreadyNode(key string) error {
	return fmt.Errorf("node is already existed: [%s]", key)
}

func NotExistNode(key string) error {
	return fmt.Errorf("node not exist: [%s]", key)
}
