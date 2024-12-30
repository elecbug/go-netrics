package graph

import "fmt"

type Identifier uint

func (id Identifier) String() string {
	return fmt.Sprintf("%d", id)
}
