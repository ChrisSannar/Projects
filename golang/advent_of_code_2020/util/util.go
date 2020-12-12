package util

import (
	"io/ioutil"
	"strconv"
)

func ReadFileToString(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	} else {
		text := string(content)
		return text;
	}
}

func ConvertStringArrayToInt(strArr []string) []int {
	nums := make([]int, len(strArr))
	for i, value := range strArr {
		num, _ := strconv.Atoi(value)
		nums[i] = num
	}
	return nums
}