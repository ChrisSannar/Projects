package datastructures

import (
	"fmt"
	"regexp"
	"strconv"
	"math"
)

// Edge is a graph edge that supports weights
type Edge struct {
	weight int
	from   *Node
	to     *Node
}

// Node is a graph node that supports weights and edges
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
	return findDepthSearchRec(value, graph.head)
}

// A recursive function for finding a node using depthfirst search
func findDepthSearchRec(value string, node *Node) (*Node, string) {

	// If we find what we're looking for, then send it up
	if node.value == value {
		return node, ""
	}

	// Make sure it's listed that we've visited this node before
	node.weight = int(math.Inf(-1))

	// if not, recursively call each edge
	for _, edge := range node.edges {
		
		// Make sure we aren't visiting a node that's already been visited
		if edge.to == nil || edge.to.weight == int(math.Inf(-1)) {
			continue
		}

		// Return if we do find it
		node, err := findDepthSearchRec(value, edge.to)
		if node != nil {
			return node, err
		}
	}

	return nil, "Unable to find value"
}

// Searches thorough a graph, recursively breadth first, to find a Node given a string value
func (graph *Graph) findBreadthSearch(value string) (*Node, string) {
	return findBreadthSearchRec(value, graph.head)
	// return &Node{edges: []*Edge{}, value: "", weight: 0}, ""
}

// A recursive function for finding a node using depthfirst search
func findBreadthSearchRec(value string, node *Node) (*Node, string) {

	node.weight = int(math.Inf(-1))

	return nil, "Unable to find value"
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
		split := regex.Split(edgeStr, int(math.Inf(-1)))
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
