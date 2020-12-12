package main

import (
	"fmt"
	"strings"
	"../util"
)

func main() {
	text := util.ReadFileToString("./input.txt")
	matrix := parseInputToMatrix(text)
	count := countSlopeEncounters(matrix, 1, 1)
	count *= countSlopeEncounters(matrix, 3, 1)
	count *= countSlopeEncounters(matrix, 5, 1)
	count *= countSlopeEncounters(matrix, 7, 1)
	count *= countSlopeEncounters(matrix, 1, 2)
	printMatrix(matrix)
	fmt.Println("Count:", count)
}

func parseInputToMatrix(text string) [][]string {
	lines := strings.Split(text, "\n")
	ylen, xlen := len(lines), len(lines[0])
	matrix := make([][]string, ylen)
	for i := range matrix {
		matrix[i] = make([]string, xlen)
	}

	// for i, line := range lines {
	for i, line := range lines {
		for j := range line {
			matrix[i][j] = string(line[j])
		}
	}

	return matrix
}

func countSlopeEncounters(matrix [][]string, right int, down int) int {
	x, y, count := 0, 0, 0
	for {
		if y >= len(matrix) {
			break;
		}

		// fmt.Println(y, x)
		if matrix[y][x] == "#" {
			// matrix[y][x] = "X"
			count++
		} 
		// else {
		// 	matrix[y][x] = "O"
		// }

		x = (x + right) % (len(matrix[0]))
		y += down
	}
	return count
}

func printMatrix(matrix [][]string) {
	fmt.Println("----------------------")
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Printf("%v ", matrix[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Println("----------------------")
	// fmt.Println(matrix)
}