package main

import (
	"fmt"
	"strings"
	"regexp"
	"../util"
)

func main() {
	text := util.ReadFileToString("./input.txt")
	groups := parseGroups(text)
	
	var count int = 0
	for _, mappers := range groups {
		// fmt.Println(mappers)
		// fmt.Println("---")
		// count += len(mappers)
		for key, value := range mappers {
			if key != "size" {
				if value == mappers["size"] {
					count++
				}
			}
		}
	}
	fmt.Println("Count", count)
}

func parseGroups(textInput string) []map[string]int {
	split := strings.Split(textInput, "\n\n")
	var groups []map[string]int = make([]map[string]int, len(split))
	whitespaceRegex := regexp.MustCompile("[^\\s]+")

	for i, val := range split {
		whitespaceSplit := whitespaceRegex.FindAllString(val, -1)
		groups[i] = make(map[string]int)
		groups[i]["size"] = len(whitespaceSplit)

		for _, val2 := range whitespaceSplit {
			for _, mark := range val2 {
				if groups[i][string(mark)] == 0 {
					groups[i][string(mark)] = 1
				}	else {
					groups[i][string(mark)]++
				}
			}
		}
	}
	return groups
}
// func parseGroups(textInput string) []map[string]bool {
// 	split := strings.Split(textInput, "\n\n")
// 	var groups []map[string]bool = make([]map[string]bool, len(split))
// 	whitespaceRegex := regexp.MustCompile("[^\\s]+")

// 	for i, val := range split {
// 		groups[i] = make(map[string]bool)
// 		whitespaceSplit := whitespaceRegex.FindAllString(val, -1)

// 		for _, val2 := range whitespaceSplit {
// 			for _, mark := range val2 {
// 				groups[i][string(mark)] = true
// 			}
// 		}
// 	}
// 	return groups
// }