package main

import (
	"fmt"
)

func main() {
	fmt.Println("WORKING")
	testAddStructToArrayStructProperty()
}

type B struct {
	value string
}

type A struct {
	arr []*B
}

func testAddStructToArrayStructProperty() {
	AMap := map[string]*A{}
	a := &A{arr: []*B{}}
	AMap["A"] = a
	bs := []*B{&B{value: "test1"}, &B{value: "test2"}, &B{value: "test3"}}
	for _, b := range bs {
		mappers := AMap["A"]
		mappers.arr = append(a.arr, b)
	}
	fmt.Println(AMap["A"])
}
