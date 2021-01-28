package datastructures

import (
	"testing"
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
	}

	if result.value != "D" {
		t.Errorf(`
Failed to find item with depth-first search
Expected: %v,
Result: %v`, "D", result.value)
	}

	result, err = graphers.findBreathSearch("D")
	if err != "" {
		t.Errorf(err)
	}

	if result.value != "D" {
		t.Errorf(`
Failed to find item with depth-first search
Expected: %v,
Result: %v`, "D", result.value)
	}
	graphers.print()
}
