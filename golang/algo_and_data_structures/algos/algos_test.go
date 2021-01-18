package algos

import (
	"testing" 
	// "fmt"
)

func TestBinarySearch(t *testing.T) {

	// Search for an existing item
	arr := []int{ 4, 8, 12, 19, 22, 47 }
	testBinarySearchSetup(t, arr, 12, 2)

	// Test for a non-existing item
	testBinarySearchSetup(t, arr, 1, -1)
	
	// Edge cases
	testBinarySearchSetup(t, arr, 4, 0)
	testBinarySearchSetup(t, arr, 47, len(arr) - 1)

	// Test empty array
	arr = []int{}
	testBinarySearchSetup(t, arr, 0, -1)
}

func TestBubbleStort(t *testing.T) {
	testSortingAlgorithm(t, BubbleSort)
}

func TestSelectionSort(t *testing.T) {
	testSortingAlgorithm(t, SelectionSort)
}

func TestInsertionSort(t *testing.T) {
	testSortingAlgorithm(t, InsertionSort)
}

func TestMergeSort(t *testing.T) {
	testSortingAlgorithm(t, MergeSort)
	// MergeSort([]int{9, 2, 3, 1, 5, 4, 7})
}

func TestQuickSort(t *testing.T) {
	testSortingAlgorithm(t, QuickSort)
	// fmt.Println(QuickSort([]int{9, 2, 3, 1, 5, 4, 7}))
}

func testBinarySearchSetup(t *testing.T, arr []int, item int, correctIndex int) {
	index := BinarySearch(arr, item)
	if index != correctIndex {
		t.Fatalf(`
Searched Item: %v
Array: %v
Found: %v
Expected: %v
`, item, arr, index, correctIndex)
	}
}

type algorithm func([]int) []int

func testSortingAlgorithm(t *testing.T, fn algorithm) {
	arr := GenerateRandomArray(8)
	sorted := fn(arr)
	if !testIfSortedArray(sorted) {
		t.Fatalf(`
Array failed to sort 
Original: %v
Sorted: %v
`, arr, sorted)
	}

	arr = []int{}
	sorted = fn(arr)
	if !testIfSortedArray(sorted) {
		t.Fatalf(`
Empty Array failed to sort`)
	}

	// Test edge cases
	arr = GenerateRandomArray(1)
	sorted = fn(arr)
	if !testIfSortedArray(sorted) {
		t.Fatalf(`
One item array failed to sort
Original: %v
Sorted: %v
`, arr, sorted)
	}

	// Odd Sorting
	arr = GenerateRandomArray(23)
	sorted = fn(arr)
	if !testIfSortedArray(sorted) {
		t.Fatalf(`
Odd Array failed to sort
Original: %v
Sorted: %v
`, arr, sorted)
	}
	
}

func testIfSortedArray(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i - 1] > arr[i] {
			return false
		}
	}
	return true
}
