# ChainedGraphs: A Scalable Graph Processing Library in Go

ChainedGraphs is a Go library designed for efficient and scalable processing of large graphs. It leverages a novel chained graph representation to minimize memory footprint and optimize traversal operations, making it suitable for applications ranging from social network analysis to recommendation systems.

This library addresses the limitations of traditional adjacency matrix and adjacency list representations when dealing with extremely large graphs. These conventional approaches often consume excessive memory, especially for sparse graphs. ChainedGraphs overcomes this issue by representing the graph as a series of smaller, interconnected subgraphs or "chains." Each chain efficiently stores a subset of the graph's edges, and these chains can be dynamically loaded and processed as needed. This allows for processing graphs that are significantly larger than available RAM, opening up new possibilities for analyzing and manipulating complex datasets. The library is built with concurrency in mind, allowing for parallel processing of graph chains to accelerate computations.

The core principle behind ChainedGraphs is locality of reference. By organizing the graph into chains, related nodes and edges are grouped together, reducing the need for random memory access during traversal operations. This improves cache utilization and significantly boosts performance. Furthermore, the chained representation facilitates efficient graph partitioning and distribution across multiple machines, paving the way for distributed graph processing in the future. ChainedGraphs provides a flexible and powerful foundation for building high-performance graph applications.

Beyond its efficient representation, ChainedGraphs offers a rich set of functionalities for graph manipulation and analysis. This includes support for various graph algorithms, such as depth-first search (DFS), breadth-first search (BFS), and shortest path algorithms like Dijkstra's and A*. The library also provides tools for graph transformation, such as subgraph extraction and node/edge filtering. Moreover, ChainedGraphs is designed to be extensible, allowing users to easily implement custom graph algorithms and data structures tailored to their specific needs. The library aims to be a versatile and adaptable solution for tackling a wide range of graph-related problems.

## Key Features

*   **Chained Graph Representation:** Efficiently stores large graphs by breaking them into interconnected chains, minimizing memory footprint and improving locality of reference. Chains are dynamically loaded/unloaded to handle graphs larger than RAM.
*   **Concurrent Processing:** Utilizes Go's concurrency primitives (goroutines and channels) to enable parallel processing of graph chains, significantly accelerating graph algorithms. The degree of concurrency is configurable.
*   **Traversal Algorithms:** Implements core graph traversal algorithms including Depth-First Search (DFS), Breadth-First Search (BFS), and Dijkstra's shortest path algorithm. These algorithms are optimized for the chained graph representation.
*   **Subgraph Extraction:** Allows for efficient extraction of subgraphs based on node or edge criteria. The extraction process preserves the chained graph structure where possible, maintaining performance.
*   **Extensible API:** Provides a flexible API for implementing custom graph algorithms and data structures. Users can easily extend the library to meet their specific requirements.
*   **Graph Partitioning (Future):** Designed to facilitate graph partitioning for distributed processing across multiple machines (planned feature for future releases).
*   **Node and Edge Attributes:** Supports associating arbitrary attributes with nodes and edges, enabling richer graph representations and more complex analysis. Attributes are stored separately to optimize memory usage for graphs with minimal metadata.

## Technology Stack

*   **Go (Golang):** The primary programming language, chosen for its performance, concurrency support, and memory management capabilities. Go's static typing and garbage collection contribute to code stability and maintainability.
*   **Go Standard Library:** Extensive use of the Go standard library for core functionalities like data structures, concurrency, and I/O operations.
*   **Testing Framework (Go's `testing` package):** Used for comprehensive unit and integration testing to ensure code quality and reliability.
*   **(Optional) Graphviz (for Visualization):** While not a direct dependency, Graphviz can be used to visualize the graph structure. The library includes utilities to export graph data in a format compatible with Graphviz.

## Installation

1.  **Install Go:** Ensure you have Go installed on your system. You can download the latest version from the official Go website: https://go.dev/dl/. Verify the installation by running `go version` in your terminal.

2.  **Clone the Repository:** Clone the ChainedGraphs repository from GitHub:
    git clone https://github.com/jjfhwang/ChainedGraphs.git

3.  **Navigate to the Project Directory:**
    cd ChainedGraphs

4.  **Build the Project:** Use the `go build` command to compile the library:
    go build .

5.  **Install the Library:** Install the library into your Go workspace using the `go install` command:
    go install .

## Configuration

ChainedGraphs uses environment variables and configuration files for customizing its behavior. The following environment variables are supported:

*   `CHAINEDGRAPHS_CHAIN_SIZE`: Specifies the maximum number of nodes per chain. A smaller chain size results in more chains but potentially better memory locality. Default value: 1000.
*   `CHAINEDGRAPHS_CONCURRENCY`: Sets the maximum number of concurrent goroutines used for graph processing. Increasing concurrency can improve performance on multi-core systems. Default value: Number of CPUs.

Configuration can also be set programmatically via the `Config` struct within the ChainedGraphs package. This allows for more fine-grained control over the library's behavior.
For example:
type Config struct {
   ChainSize int
   Concurrency int
}

## Usage

Here's a basic example of how to use ChainedGraphs:

// Create a new chained graph
graph := NewChainedGraph()

// Add nodes to the graph
graph.AddNode("A")
graph.AddNode("B")
graph.AddNode("C")

// Add edges to the graph
graph.AddEdge("A", "B")
graph.AddEdge("B", "C")
graph.AddEdge("C", "A")

// Perform a Depth-First Search starting from node "A"
result := graph.DFS("A")

// Print the result
fmt.Println(result)

For detailed API documentation, please refer to the generated GoDoc documentation. Run `go doc` within the project directory to access the documentation. Example: `go doc ChainedGraphs.ChainedGraph`

## Contributing

We welcome contributions to ChainedGraphs! Please follow these guidelines:

1.  Fork the repository.
2.  Create a new branch for your feature or bug fix.
3.  Write clear and concise commit messages.
4.  Include unit tests for your changes.
5.  Submit a pull request with a detailed description of your changes.
6. All new code should adhere to the Go coding standards (gofmt, golint).

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/jjfhwang/ChainedGraphs/blob/main/LICENSE) file for details.

## Acknowledgements

We would like to thank the Go community for providing excellent tools and resources that made this project possible.