package algos

// BinarySearch searches for a given item inside an array and returns the index of that given item
func BinarySearch(arr []int, item int) int {
	var high, low int = len(arr) - 1, 0
	for {
		if low > high {
			return -1
		}

		mid := (high + low) / 2
		if arr[mid] < item {
			low = mid + 1

		} else if arr[mid] > item {
			high = mid - 1
		
		} else {
			return mid
		}
	}
}

// BubbleSort sorts an array using bubblesort
func BubbleSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	for i := 1; i < len(arr); i++ {
		index1, index2 := 0, 1
		for index2 < len(arr) {
			if arr[index1] > arr[index2] {
				swap(arr, index1, index2)
			}
			index1++
			index2++
		}
	}
	return arr
}

// SelectionSort sorts a given array using selectionsort
func SelectionSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	for i := range arr {
		min, j := i, i + 1
		for ;j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		swap(arr, i, min)
	}

	return arr
}

// InsertionSort sorts an array with insertionsort
func InsertionSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	for i := 1; i < len(arr); i++ {
		j := i 
		for j - 1 >= 0 && arr[j] < arr[j - 1] {
			swap(arr, j, j - 1)
			j--
		}
	}

	return arr
}

// MergeSort sorts an array with mergesort
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	midIndex := len(arr) / 2
	left := MergeSort(arr[:midIndex])
	right := MergeSort(arr[midIndex:])
	return merge(left, right)

}

// Merges two arrays together in order
func merge(arr1 []int, arr2 []int) []int {

	// Create the resulting array and assign the indicies
	result := make([]int, len(arr1) + len(arr2))
	resultIndex, index1, index2 := 0, 0, 0
	for index1 < len(arr1) && index2 < len(arr2) {

		// If arr1 has the lower number, insert that in and vice versa
		if arr1[index1] <= arr2[index2] {
			result[resultIndex] = arr1[index1]
			index1++
		} else {
			result[resultIndex] = arr2[index2]
			index2++
		}
		resultIndex++
	}

	// Copy the remaining elements over
	for index1 < len(arr1) {
		result[resultIndex] = arr1[index1]
		index1++
		resultIndex++
	}
	for index2 < len(arr2) {
		result[resultIndex] = arr2[index2]
		index2++
		resultIndex++
	}

	return result
}

// QuickSort sorts an array with quicksort
func QuickSort(arr []int) []int {
	return []int{}
}
