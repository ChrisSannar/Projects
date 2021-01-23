package datastructures

import "fmt"

type Edge struct {
	weight int
	from   Node
	to     Node
}

type Node struct {
	edges  []Edge
	value  string
	weight int
}

type Graph struct {
	size  int
	nodes map[string]Node
}

func main() {
	fmt.Println("Working")
}

func CreateWeightedGraphFromInput(nodes []string, edges []string) Graph {
	return Graph{
		size:  0,
		nodes: make(map[string]Node),
	}
}
