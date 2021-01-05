package main

import (
	"fmt"
	"strings"

	"../util"
)

const preambleCount int = 25
const del string = "\n"

func main() {
	text := util.ReadFileToString("./input.txt")
	nums := util.ConvertStringArrayToInt(strings.Split(text, del))

	var badNumIndex int = findBadNumberIndex(nums, preambleCount)
	fmt.Println("Bad Number:", nums[badNumIndex])

	var seq []int = findSummedSequence(nums, badNumIndex)
	smallest := findSmallestNumber(seq)
	largest := findLargestNumber(seq)
	fmt.Println("Sequence", smallest+largest)
}

func findBadNumberIndex(nums []int, pre int) int {
	for i := pre; i < len(nums); i++ {
		if !containsSummedNumbers(nums[(i-pre):i], nums[i]) {
			return i
		}
	}
	return -1
}

func containsSummedNumbers(nums []int, target int) bool {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return true
			}
		}
	}
	return false
}

func findSummedSequence(nums []int, targetIndex int) []int {
	for i := 0; i < targetIndex; i++ {
		sum := 0
		for j := i; j < targetIndex; j++ {
			sum += nums[j]
			if sum == nums[targetIndex] {
				return nums[i : j+1]
			}
		}
	}
	return []int{}
}

func findSmallestNumber(nums []int) int {
	smallest := int(^uint(0) >> 1)
	for _, num := range nums {
		if num < smallest {
			smallest = num
		}
	}
	return smallest
}

func findLargestNumber(nums []int) int {
	largest := -1
	for _, num := range nums {
		if num > largest {
			largest = num
		}
	}
	return largest
}
