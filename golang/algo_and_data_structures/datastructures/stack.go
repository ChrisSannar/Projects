package datastructures

// import (
// 	"fmt"
// )

type Stack struct {
	size int
	arr	[]int
}

func CreateStack() *Stack {
	return &Stack {size: 0, arr: []int{}}
} 