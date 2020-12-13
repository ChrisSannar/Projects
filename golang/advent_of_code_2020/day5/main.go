package main

import (
	"fmt"
	"strings"
	"../util"
)

func main() {
	text := util.ReadFileToString("./input.txt")
	split := strings.Split(text, "\n")
	// var highest int = 0
	// var mappers map[int]bool = make(map[int]bool)
	var arr [1024]int
	for _, val := range split {
		// fmt.Println(val, len(val))
		row, column := getRowAndColumnValues(val)
		var index int = (row * 8) + column
		arr[index] = 1
		// if arr[arr[index] + 1] > 0 && arr[arr[index] - 1] > 0 {
		// 	// if (row * 8) + column > highest {
		// 	// 	highest = (row * 8) + column
		// 	// }
		// 	fmt.Println(index)
		// }

	}
	for i := 1; i < len(arr) - 1; i++ {
		if (arr[i] == 0 && arr[i + 1] > 0 && arr[i - 1] > 0) {
			fmt.Println(i)
		}
	}
	// fmt.Println("Highest:", highest)
}

func getRowAndColumnValues(textInput string) (int, int) {
	rowSearch := binarySearch(textInput[:7], 127, "F")
	columnSearch := binarySearch(textInput[7:], 7, "L")
	
	return rowSearch, columnSearch
}

func binarySearch(input string, limit int, lowerChar string) int {
	low, high := 0, limit
	for _, char := range input {
		if string(char) == lowerChar {
			high = (low + high) / 2
		} else {
			low = (low + high) / 2 + 1
		}
	}
	return (low + high) / 2
}

