package main

import (
	"fmt"
	"sort"
	"strings"

	"../util"
)

const del = "\r\n"

func main() {
	text := util.ReadFileToString("./test3.txt")
	nums := util.ConvertStringArrayToInt(strings.Split(text, del))
	oneJoltCount, threeJoltCount := countDifferingJoltage(nums)

	fmt.Println(oneJoltCount, "*", threeJoltCount, "=", oneJoltCount*threeJoltCount)

	fmt.Println("Count:", countPossibleArangements(nums))
}

func countDifferingJoltage(nums []int) (int, int) {

	// First sort the nums
	sort.Ints(nums)

	// Next, count if we have a 1 jolt difference or a 3 jolt difference
	oneJoltCount, threeJoltCount := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] == 1 {
			oneJoltCount++
		} else if nums[i+1]-nums[i] == 3 {
			threeJoltCount++
		}
	}
	return oneJoltCount + 1, threeJoltCount + 1
}

func countPossibleArangements(nums []int) int {

	// Sort out nums
	sort.Ints(nums)

	// We're going to have to iterate over each each element to see the possibilities
	count := 0
	j := 1
	for i := 0; i < len(nums)-1; i++ {
		fmt.Println("LOOP")
	nextNumber:
		currentNumber, nextNumber := nums[i], nums[j]
		fmt.Println(currentNumber, nextNumber)
		if nextNumber-currentNumber > 3 {
			fmt.Println("MORE")
			continue
		} else if nextNumber-currentNumber == 3 {
			fmt.Println("THREE")
			// count++
			j++
			continue
		} else if nextNumber-currentNumber == 2 {
			fmt.Println("TWO")
			count++
			j++
			goto nextNumber
		} else {
			fmt.Println("ONE")
			count++
			j++
			goto nextNumber
		}
	}
	return count
}

// 1 2 3 4 7 10 <-- largest
// 1 2 4 7 10
// 1 3 4 7 10
// 1 4 7 10 <-- smallest
