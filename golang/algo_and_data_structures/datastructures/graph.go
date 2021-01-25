package datastructures

import (
	"fmt"
	"regexp"
	"strconv"
)

type Edge struct {
	weight int
	from   Node
	to     Node
}

type Node struct {
	edges  []*Edge
	value  string
	weight int
}

type Graph struct {
	size  int
	nodes map[string]*Node
}

func (graph *Graph) print() {
	for _, node := range graph.nodes {
		fmt.Println(node)
	}
}

func appendEdge(node Node, edge *Edge) []*Edge {
	edges := node.edges
	edges = append(edges, edge)
	return edges
}

func CreateWeightedGraphFromInput(nodes []string, edges []string) Graph {
	nodeMap := make(map[string]*Node)
	size := 0

	// Set all the nodes in the graph with their default weights of -1
	for _, nodeStr := range nodes {
		nodeMap[nodeStr] = &Node{
			edges:  []*Edge{},
			value:  nodeStr,
			weight: -1,
		}
		size++
	}

	// Include all the edges
	regex := regexp.MustCompile("\\s*(\\(|\\))\\s*") // Proper regex for parsing the input
	for _, edgeStr := range edges {

		// Split the values out and parse them into an edge
		split := regex.Split(edgeStr, -1)
		weight, _ := strconv.Atoi(split[1])
		edge := &Edge{
			weight: weight,
			from:   *nodeMap[split[0]],
			to:     *nodeMap[split[2]],
		}

		// Then insert that edge into the given Node
		noders := nodeMap[split[0]]
		noders.edges = append(noders.edges, edge)
	}

	// Resulting graph object
	return Graph{
		size:  size,
		nodes: nodeMap,
	}
}
