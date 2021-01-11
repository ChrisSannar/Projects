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
	return []int{}
}


// QuickSort sorts an array with quicksort
func QuickSort(arr []int) []int {
	return []int{}
}
