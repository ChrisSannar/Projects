package datastructures

import (
	"testing" 
	// "fmt"
)

func TestDoubleLinkedList(t *testing.T) {
	list := CreateDoubleLinkedList()
	testDoubleLinkedListInsert(t, &list)
	testDoubleLinkedListGet(t, &list)
	testDoubleLinkedListRemove(t, &list)
}

func testDoubleLinkedListGet(t *testing.T, list *DoubleLinkedList) {
	val := list.get(0).value
	if val != 2 {
		t.Fatalf("Failed to get first value of list\n")
	}
	
	val = list.get(2).value
	if val != 9 {
		t.Fatalf("Failed to get last value of list\n")
	}

	val = list.get(1).value
	if val != 3 {
		t.Fatalf("Failed to get middle value of list\n")
	}
}

func testDoubleLinkedListInsert(t *testing.T, list *DoubleLinkedList) {
	
	// Insert out of bounds
	err := list.insert(5, 1)
	if err == "" {
		t.Fatalf("Failed to catch out of bounds index\n")
	}
	err = list.insert(5, -1)
	if err == "" {
		t.Fatalf("Failed to catch out of bounds index\n")
	}

	// Insert at the beginning
	err = list.insert(2, 0)
	if err != "" {
		t.Fatalf(err)
	}
	
	// Insert at the end
	err = list.insert(9, 1)
	if err != "" {
		t.Fatalf(err)
	}

	// Insert in the middle
	err = list.insert(3, 1)
	if err != "" {
		t.Fatalf(err)
	}

	if list.size != 3 {
		t.Fatalf("List size off\nExpected: %v\nSize: %v\n", 3, list.size)
	}

	// Check if we have the correct values
	if list.get(0).value != 2 || 
		list.get(1).value != 3 || 
		list.get(2).value != 9 {
		t.Fatalf("Improper list order: \nExpected: %v\nList:\t  {2} <-> {3} <-> {9}\n", list.toString())
	}
}

func testDoubleLinkedListRemove(t *testing.T, list *DoubleLinkedList) {
	
	// Check for out of bounds
	err := list.remove(4)
	if err == "" {
		t.Fatalf("Failed to catch out of bounds index\n")
	}
	err = list.remove(-1)
	if err == "" {
		t.Fatalf("Failed to catch out of bounds index\n")
	}
	
	// Check for proper removal
	err = list.remove(1)
	if err != "" && list.get(0).value == 2 && list.get(1).value == 9 {
		t.Fatalf("Failed to remove middle element")
	}

	err = list.remove(1)
	if err != "" && list.get(0).value == 2 {
		t.Fatalf("Failed to remove last element")
	}

	err = list.remove(0)
	if err != "" {
		t.Fatalf("Failed to remove first element")
	}

	if list.get(0) != nil {
		t.Fatalf("List failed to remove items: %v", list.toString())
	}

	if list.size > 0 {
		t.Fatalf("Failed to update list size")
	}

}