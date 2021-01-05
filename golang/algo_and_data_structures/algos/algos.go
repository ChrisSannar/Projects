package algos

import "fmt"

// SelectionSort sorts a given array using selection sort
func SelectionSort(arr []int) {
	swap(arr, 0, 1)
}

// BinarySearch searches for a given item inside an array and returns the index of that given item
func BinarySearch(arr []int, item int) int {
	var high, low int = len(arr) - 1, 0
	for {
		mid := (high + low) / 2
		if arr[mid] < item {
			low = mid + 1

		} else if arr[mid] > item {
			high = mid - 1
		
		} else {
			return mid
		}

		if low > high {
			return -1
		}
	}
}
