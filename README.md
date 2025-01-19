# Go-netrics: Advanced Graph Algorithm Framework

Go-netrics is a high-performance graph algorithm framework written in Go. It provides efficient tools for graph construction, traversal, and advanced computations such as shortest paths, centrality measures, and graph transformations. The framework supports both sequential and parallel computation modes, making it suitable for handling graphs of various sizes and complexities.

---

## Features

- **Graph Representation**
  - Supports directed, undirected, weighted, and unweighted graphs.
  - Provides adjacency matrix and adjacency list representations.

- **Algorithms**
  - Shortest path computations and more metrics computations:
    - Weighted graphs: Dijkstra's algorithm.
    - Unweighted graphs: Breadth-First Search (BFS).
  - Parallelized computations for large graphs.

- **Modular Design**
  - `Unit`: Sequential computation unit for small to medium-sized graphs.
  - `ParallelUnit`: Parallel computation unit leveraging multiple CPU cores.

- **Extensible**
  - Flexible API for adding custom algorithms.
  - Clean structure for graph manipulations.

---

## Installation

Ensure you have Go installed (version 1.16 or later). Then run:

```bash
go get -u github.com/elecbug/go-netrics
```

---

## Usage

```go
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/davecgh/go-spew/spew"
	netrics "github.com/elecbug/go-netrics"
)

func main() {
    cap := 200
	g := netrics.NewGraph(netrics.UNDIRECTED_UNWEIGHTED, cap)

	for i := 0; i < cap; i++ {
		g.AddNode(fmt.Sprintf("%4d", i))
	}

	for i := 0; i < g.NodeCount()*g.NodeCount()/10; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(i)))
		from := netrics.NodeID(r.Intn(g.NodeCount()))

		r = rand.New(rand.NewSource(time.Now().UnixNano() + int64(i*i)))
		to := netrics.NodeID(r.Intn(g.NodeCount()))

		g.AddEdge(from, to)
	}

	fmt.Printf("\n%s\n", g.String())

	{
		u := g.ToUnit()
		s := time.Now()

		fmt.Printf("\nShortestPath: %v\n", spew.Sdump(u.ShortestPath(0, netrics.NodeID(cap-1))))
		fmt.Printf("\nAverageShortestPathLength: %v\n", spew.Sdump(u.AverageShortestPathLength()))
		fmt.Printf("\nBetweennessCentrality: %v\n", spew.Sdump(u.BetweennessCentrality()))
		fmt.Printf("\nClusteringCoefficient: %v\n", spew.Sdump(u.ClusteringCoefficient()))
		fmt.Printf("\nDegreeCentrality: %v\n", spew.Sdump(u.DegreeCentrality()))
		fmt.Printf("\nDiameter: %v\n", spew.Sdump(u.Diameter()))
		fmt.Printf("\nEigenvectorCentrality: %v\n", spew.Sdump(u.EigenvectorCentrality(1000, 1e-6)))
		fmt.Printf("\nGlobalEfficiency: %v\n", spew.Sdump(u.GlobalEfficiency()))
		fmt.Printf("\nLocalEfficiency: %v\n", spew.Sdump(u.LocalEfficiency()))
		fmt.Printf("\nPercentileShortestPathLength: %v\n", spew.Sdump(u.PercentileShortestPathLength(30)))
		fmt.Printf("\nRichClubCoefficient: %v\n", spew.Sdump(u.RichClubCoefficient(2)))

		duration := time.Since(s)
		fmt.Printf("execution time: %s", duration)
	}
	{
		pu := g.ToParallelUnit(20)
		s := time.Now()

		fmt.Printf("\nShortestPath: %v\n", spew.Sdump(pu.ShortestPath(0, netrics.NodeID(cap-1))))
		fmt.Printf("\nAverageShortestPathLength: %v\n", spew.Sdump(pu.AverageShortestPathLength()))
		fmt.Printf("\nBetweennessCentrality: %v\n", spew.Sdump(pu.BetweennessCentrality()))
		fmt.Printf("\nClusteringCoefficient: %v\n", spew.Sdump(pu.ClusteringCoefficient()))
		fmt.Printf("\nDegreeCentrality: %v\n", spew.Sdump(pu.DegreeCentrality()))
		fmt.Printf("\nDiameter: %v\n", spew.Sdump(pu.Diameter()))
		fmt.Printf("\nEigenvectorCentrality: %v\n", spew.Sdump(pu.EigenvectorCentrality(1000, 1e-6)))
		fmt.Printf("\nGlobalEfficiency: %v\n", spew.Sdump(pu.GlobalEfficiency()))
		fmt.Printf("\nLocalEfficiency: %v\n", spew.Sdump(pu.LocalEfficiency()))
		fmt.Printf("\nPercentileShortestPathLength: %v\n", spew.Sdump(pu.PercentileShortestPathLength(30)))
		fmt.Printf("\nRichClubCoefficient: %v\n", spew.Sdump(pu.RichClubCoefficient(2)))

		duration := time.Since(s)
		fmt.Printf("execution time: %s", duration)
	}
}
```

---

## Contributing

Go-netrics is an open-source project, and contributions are welcome! Whether it's fixing bugs, improving documentation, or adding new features, your input can make a difference.

### How to Contribute

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/my-feature`).
3. Make your changes and test thoroughly.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push the branch to your fork (`git push origin feature/my-feature`).
6. Open a pull request.

Please ensure your contributions adhere to the project's coding guidelines and include relevant tests.

---

## License

Go-netrics is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

---

## Contact

If you have any questions, feedback, or feature requests, feel free to open an issue on GitHub or contact us via email at [Email](mailto:deveb1479@gmail.com).