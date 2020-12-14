package main

import (
	"fmt"
	"strings"
	"strconv"
	"../util"
)

const lineDelimiter string = ".\n"

func main() {
	text := util.ReadFileToString("./input.txt")
	bags := mapBagString(text)
	
	matrix := buildBagMatrix(bags)
	// printMatrix(matrix)
	// fmt.Println()

	const searchedBag string = "shiny gold bag"
	
	var count int = 0;

	// Check each of the bags
	for _, containing := range matrix {
		// fmt.Println("BAG", bag)

		// For each of those bags, check each of the containing ones
		for baggers, _ := range containing {
			// fmt.Println("CONTAIN", baggers)

			// If it directly contains the bag or contains a bag that does, count it
			if  strings.Contains(baggers, searchedBag) || canContainBag(matrix, baggers, searchedBag) {
				count++
				break;	// Once we know the bag contains it, we're done
			}
		}
		// fmt.Println()
	}
	fmt.Println("Count:", count)

	// Minus one to exclude the bag itself
	fmt.Println("Count2:", countContainingBags(matrix, searchedBag) - 1)
	/*
	/* */
}

func countContainingBags(matrix map[string]map[string]int, bagName string) int {
	// fmt.Println("NAME:", bagName)
	
	// Base case
	if len(matrix[bagName]) == 0{
		return 1
	} else {
		var total int = 1;
		for bag, count := range matrix[bagName] {
			// fmt.Println(bag, count)
			total += (countContainingBags(matrix, bag) * count)
		}
		// fmt.Println("TOTAL", bagName, total)
		// fmt.Println()
		return total
	}
}


func canContainBag(matrix map[string]map[string]int, bagName string, searchedBag string) bool {
	// fmt.Println("-", bagName, "-")

	// If our current bag doesn't have any containing, then it won't be in there
	if len(matrix[bagName]) == 0 {
		return false
	}

	// Look through each of the congaining bags the passed in bag has
	for baggers, _ := range matrix[bagName] {
		// fmt.Println("SEARCHING", baggers)
		if strings.Contains(baggers, searchedBag) {
			// fmt.Println("FOUND")
			return true
		} else {
			// fmt.Println("RECURSE")
			if canContainBag(matrix, baggers, searchedBag) {
				return true
			}
		}
		// fmt.Println("--- NEXT BAG ---")
	}
	return false
}

func mapBagString(textInput string) map[string]string {
	lines := strings.Split(textInput, lineDelimiter)
	// fmt.Println(len(lines))

	// Doing this to chop off the period on the last one
	lastLine := lines[len(lines) - 1]
	lines[len(lines) - 1] = lastLine[:len(lastLine) - 1]
	
	// Set each line to it's mapped value
	var result map[string]string = make(map[string]string)
	for _, line := range lines {
		split := strings.Split(line, " contain ")

		result[eliminatePlural(split[0])] = split[1]
	}
	return result
}

// Builds out the matrix based on the matrix bag already
func buildBagMatrix(stringMap map[string]string) map[string]map[string]int {
	
	// Our starting matrix 
	var matrix map[string]map[string]int = make(map[string]map[string]int, len(stringMap))
	
	for key, value := range stringMap {
		matrix[key] = make(map[string]int)
		for name, count := range parseBagsFromString(value) {
			// fmt.Print(name, count)
			matrix[key][eliminatePlural(name)] = count
		}
	}
	return matrix
}

// Creates a map of bags and their counts from an incoming string
func parseBagsFromString(bagString string) map[string]int {
	var mappers map[string]int = make(map[string]int)
	bags := strings.Split(bagString, ", ")
	
	for _, bag := range bags {
		if bag != "no other bags" {
			count, _ := strconv.Atoi(string(bag[0])) 
			name := bag[2:]
			// fmt.Print(count, ":", name, ",")
			mappers[name] = count
		}
	}
	// fmt.Println()
	return mappers
}

func printMatrix(matrix map[string]map[string]int) {
	for key, val := range matrix {
		fmt.Println(key, val)
	}
}

// Just chops off the 's' on the end of a word if it's there
func eliminatePlural(word string) string {
	if word[len(word) - 1] == 's' {
		word = word[:len(word) - 1]
	}
	return word
}