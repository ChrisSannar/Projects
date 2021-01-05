package main

import (
	"fmt"
	// "strings"
	"regexp"
	"./advent_of_code_2020/util"
)

// Comment

/*

Show documentation for a package
godoc <package> <function> 
staticly typed
*/ 

func main() {

	text := util.ReadFileToString("./temp.txt")
	// lines := strings.Split(text, "\n")
	reg := regexp.MustCompile(`[^\S\r\n]+`)
	fmt.Println("---")
	fmt.Println(reg.ReplaceAllString(text, ","))
	fmt.Println("---")
	/*
	addUp := func (nums []int) (string, int) {
		sum := 0
		
		for _, value := range nums {
			sum += value
		}
		
		return "SUM", sum
	}
	
	// variable name type = value
	var asdf_1234 int = 42
	
	var num float64 = 1.68974e2

	// Auto assign number
	num2 := 14

	var numOne = 1.000
	var num99 = 0.999
	
	// Create constants
	const pi float32 = 3.14159265
	
	// + - * / %

	fmt.Printf("%.3f %.4f \n", numOne - num99, pi)
	fmt.Printf("%T %T \n", numOne - num99, pi)
	
	var (
		varA = 10
		varB = 11
	)
	
	// string -> "" ``
	
	var str string = "NAME"
	alsoStr := `I'm a string`
	
	fmt.Println(alsoStr + str)

	
	var boolers bool = true
	
	fmt.Println("This is true: ", boolers || (false && !true))
	
	for i:= 0; i <= 10; i++ {
		if i < 5 {
			fmt.Print("UNDER ")
		} else if i % 2 == 0 {
		fmt.Print("OVER5 ")
		} else {
			fmt.Print("ELSE ")
		}
		
		fmt.Print(i, " ")
	}
	
	fmt.Println()
	
	// var nums[5] int
	// nums[0] = 2
	nums := []int {5, 2, 3, 1, 4}

	for _, value := range nums {
		fmt.Print(value, ",")
	}
	fmt.Println()
	text, sum := addUp(nums)
	fmt.Println(nums[1:4], text, sum) // array slice [:2] [2:3] [3:]

	mappers := make(map[string] int)

	mappers["qwer"] = 14;
	mappers["asdf"] = 42;
	delete(mappers, "qwer")
	fmt.Println(mappers)

	fmt.Println("Hello World", asdf_1234, num, num2, factorial(varA), factorial(varB))
	/* */
}

// func factorial (num int) int {
// 	if num < 0 {
// 		return 0
// 	}

// 	if num == 0 {
// 		return 1;
// 	} else {
// 		return (factorial(num - 1) * num);
// 	}
// }

// func nextPrime(n int) int {
	
// }