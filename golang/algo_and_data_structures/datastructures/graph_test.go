package datastructures

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	graphers := CreateWeightedGraphFromInput([]string{}, []string{})
	fmt.Println("GRAPH TEST", graphers)
}
