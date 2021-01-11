package algos

import "testing"

func TestGenerateOrderedArray(t *testing.T) {
	arr := GenerateOrderedArray(0)
	if len(arr) != 0 {
		t.Fatalf(`
Should have created an empty array`)
	}

	length := 5
	arr = GenerateOrderedArray(length)
	if len(arr) != length {
		t.Fatalf(`
Length of array incorrect
	Expected: %v
	Result: %v`, length, len(arr))
	}

	for i := 1; i < len(arr); i++ {
		if arr[i - 1] > arr[i] {
			t.Fatalf(`
Array is Not in Order: %v
Failed Indicies: %v and %v`, arr, i - 1, i)
		}
	}
}

func TestGenerateRandomArray(t *testing.T) {
	arr := GenerateRandomArray(0)
	if len(arr) != 0 {
		t.Fatalf(`
Should have created an empty array`)
	}

	length := 5
	arr = GenerateRandomArray(length)
	if len(arr) != length {
		t.Fatalf(`
Length of array incorrect
	Expected: %v
	Result: %v`, length, len(arr))
	}

	mixCount := 0
	for i := 1; i < len(arr); i++ {
		if arr[i - 1] > arr[i] {
			mixCount++
		}
	}

	if mixCount < (len(arr) / 4) {
		t.Fatalf(`
Array has too much order
Size: %v
Mix Count: %v`, len(arr), mixCount)
	}

}

func TestSwap(t *testing.T) {
	arr := []int{ 0, 1 }
	swap(arr, 0, 1)
	if arr[0] != 1 && arr[1] != 0 {
		t.Fatalf(`
swap failed: %v`, arr)
	}
}
