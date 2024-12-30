package graph_err_msg

import "fmt"

func InvalidEdge(graphKey, edgeKey string) string {
	return fmt.Sprintf("Invalid edge for graph type: [(%s) does not fit %s]", edgeKey, graphKey)
}

func AlreadyNode(key string) string {
	return fmt.Sprintf("Node is already existed: [%s]", key)
}

func NotExistNode(key string) string {
	return fmt.Sprintf("Node not exist: [%s]", key)
}
