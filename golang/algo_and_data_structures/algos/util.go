package algos

import (
	"math/rand"
)

// GenerateOrderedArray returns an array of ordered values given the size
func GenerateOrderedArray(size int) []int {
	var arr []int = make([]int, size)
	for index := range arr {
		arr[index] = index
	}
	return arr
}

// GenerateRandomArray returns an array of ordered values given the size
func GenerateRandomArray(size int) []int {
	var arr []int = make([]int, size)
	for index := range arr {
		arr[index] = index
	}

	for index := range arr {
		randomNumber := rand.Intn(size)
		swap(arr, index, randomNumber)
	}

	return arr
}

// Swaps two indicies in an array
func swap(arr []int, i int, j int) {
	var temp int = arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}