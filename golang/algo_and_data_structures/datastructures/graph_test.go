package datastructures

import (
	"testing"
	"fmt"
)

func TestGraph(t *testing.T) {
	nodes := []string{"A", "B", "C", "D"}
	edges := []string{"A (3) B", "A (-2) C", "A (1) D", "B (0) A", "B (2) C", "B (4) D", "C (-1) D"}
	graphers := CreateWeightedGraphFromInput(nodes, edges)
	if len(graphers.nodes) != graphers.size {
		t.Errorf(`
Graph Nodes failed to insert
Expected: %v,
Result: %v`, graphers.size, len(graphers.nodes))
	}

	if graphers.head.value != "A" {
		t.Errorf(`
Failed to set graph head
Expected: %v
Result: %v`, "A", graphers.head.value)
	}

	result, err := graphers.findDepthSearch("D")
	if err != "" {
		t.Errorf(err)
	} else if result.value != "D" {
		t.Errorf(`
Failed to find item with depth-first search
Expected: %v,
Result: %v`, "D", result.value)
	}

	
	// Need to reset the node weights
	graphers = CreateWeightedGraphFromInput(nodes, edges)
	result, err = graphers.findDepthSearch("E")
	if err == "" {
		t.Errorf(`
Found an Node with depth first search that doesn't exist: %v`, result)
	}

	graphers = CreateWeightedGraphFromInput(nodes, edges)
	result, err = graphers.findBreadthSearch("D")
	if err != "" {
		t.Errorf(err)
	} else if result.value != "D" {
		t.Errorf(`
Failed to find item with breadth-first search
Expected: %v,
Result: %v`, "D", result.value)
	}
	
	// Need to reset the node weights
	graphers = CreateWeightedGraphFromInput(nodes, edges)
	result, err = graphers.findBreadthSearch("E")
	if err == "" {
		t.Errorf(`
Found an Node with breadth first search that doesn't exist: %v`, result)
	}
	
	start, err := graphers.findBreadthSearch("A")
	djk, err := graphers.dijkstras(start, result)
	if err != "" {
		t.Errorf(err)
	}

	djkStr := ""
	for _, n := range djk {
		djkStr = djkStr + n.value
	}
	fmt.Println(djkStr)

	if djkStr != "ACD" {
		t.Errorf(`
Incorrect Path for dijkstras algorithm
Expected: %v
Result: %v`, "ACD", djkStr)
	}

	graphers.print()
}
