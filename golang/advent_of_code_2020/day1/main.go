package main

import (
	"fmt" 
	"strings"
)
import "../util"

func main() {
	result := day1(2020)
	fmt.Println(result)
}

func day1(mark int) int {
	text := util.ReadFileToString("./input.txt")
	var split []string = strings.Split(text, "\n");
	var nums []int = util.ConvertStringArrayToInt(split);
	// fmt.Println(nums)
	
	for i, value := range nums {
		for j:= i; j < len(nums); j++ {
			// fmt.Print(nums[j], " ")
			// fmt.Print(nums[j], nums[j] + value, " ")
			// if (nums[j] + value == mark) {
			// 	return nums[j] * value
			// }
			for k:= j; k < len(nums); k++ {
				// fmt.Print(value, nums[j], nums[k], " ")
				// fmt.Print(nums[k], nums[k] + value, " ")
				if (nums[k] + nums[j] + value == mark) {
					return nums[k] * nums[j] * value
				}
			}
			// fmt.Println()
		}
		// fmt.Println()
		// fmt.Println()
	}
	return -1
}