package datastructures

import (
	"fmt"
	"regexp"
	"strconv"
)

type Edge struct {
	weight int
	from   *Node
	to     *Node
}

type Node struct {
	edges  []*Edge
	value  string
	weight int
}

// Graph is a datastructure that uses nodes to plot connections
type Graph struct {
	size  int
	nodes map[string]*Node
	head  *Node
}

func (graph *Graph) print() {
	for _, node := range graph.nodes {
		fmt.Printf(`%v (%v)[`, node.value, node.weight)
		for _, edge := range node.edges {
			fmt.Printf(` %v (%v),`, edge.to.value, edge.weight)
		}
		fmt.Println("]")
	}
}

// Searches thorough a graph, recursively depth first, to find a Node given a string value
func (graph *Graph) findDepthSearch(value string) (*Node, string) {
	return &Node{edges: []*Edge{}, value: "", weight: 0}, ""
}

// Searches thorough a graph, recursively breath first, to find a Node given a string value
func (graph *Graph) findBreathSearch(value string) (*Node, string) {
	return &Node{edges: []*Edge{}, value: "", weight: 0}, ""
}

// Finds the cheapest path between two nodes
func (graph *Graph) dijkstras(start *Node, end *Node) ([]*Node, string) {
	return []*Node{}, ""
}

func appendEdge(node *Node, edge *Edge) []*Edge {
	edges := node.edges
	edges = append(edges, edge)
	return edges
}

// CreateWeightedGraphFromInput nodes format: {"A", "B", ...}, edge format: {"A (-3) B", "A (2) C", ...}
func CreateWeightedGraphFromInput(nodes []string, edges []string) Graph {
	nodeMap := make(map[string]*Node)
	size := 0
	var head *Node

	// Set all the nodes in the graph with their default weights of -1
	for i, nodeStr := range nodes {
		nodeMap[nodeStr] = &Node{
			edges:  []*Edge{},
			value:  nodeStr,
			weight: -1,
		}
		if i == 0 {
			head = nodeMap[nodeStr]
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
			from:   nodeMap[split[0]],
			to:     nodeMap[split[2]],
		}

		// Then insert that edge into the given Node
		noders := nodeMap[split[0]]
		noders.edges = append(noders.edges, edge)
	}

	// Resulting graph object
	return Graph{
		size:  size,
		nodes: nodeMap,
		head:  head,
	}
}
