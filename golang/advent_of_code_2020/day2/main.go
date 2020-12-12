package main

import (
	"fmt"
	"strings"
	"strconv"
	"../util"
)

func main() {
	text := util.ReadFileToString("./input.txt")
	// fmt.Println("day 2", text)
	parseDatabasePasswordValues(text)
}

func parseDatabasePasswordValues(text string) {
	parseSplit := strings.Split(text, "\n")
	var count int = 0;
	for _, stringVal := range parseSplit {
		currentVals := strings.Split(stringVal, " ")
		limitString, item, password := currentVals[0], string(currentVals[1][0]), currentVals[2]
		limitStringSplit := strings.Split(limitString, "-")
		low, _ := strconv.Atoi(limitStringSplit[0])
		high, _ := strconv.Atoi(limitStringSplit[1])

		fmt.Println(low, high, item, password)
		
		// var charOccuranceCount int = 0
		// for _, charInPassword := range password {
			
		// 	if string(charInPassword) == item {
		// 		charOccuranceCount++
		// 	}
		// }

		// if charOccuranceCount >= low && charOccuranceCount <= high {
		// 	count++
		// }

		pos1, pos2 := string(password[low - 1]), string(password[high - 1])
		if (pos1 == item) != (pos2 == item) {
			count++
		}

		// fmt.Printf("%v: %v", item, charOccuranceCount)
		// fmt.Println()
	}
	fmt.Printf("Total: %v\n", count)
	// fmt.Printf("%+v\n", split)
}