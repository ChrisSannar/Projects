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
	graphers.print()
}
