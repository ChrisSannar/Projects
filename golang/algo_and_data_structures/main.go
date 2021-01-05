package main

import (
	"fmt"
	"./algos"
)

func main() {
	// arr := algos.GenerateOrderedArray(10)
	arr := []int{2, 9, 12, 33, 68}
	index := algos.BinarySearch(arr, 68)
	if index < 0 {
		fmt.Println("Item not found")
	} else {
		fmt.Println("Item found at index:", index)
	}
	// arr := algos.GenerateRandomArray(10)
	// algos.SelectionSort(arr)
	// fmt.Println("algo", arr)
}