package datastructures

import (
	"fmt"
	"strconv"
)

type listNode struct {
	value int
	next *listNode
	prev *listNode
}

// DoubleLinkedList hold integers in a list type
type DoubleLinkedList struct {
	size int
	head *listNode
	tail *listNode
}

func (dll *DoubleLinkedList) toString() string {
	result := ""
	if dll.size > 0 {
		node := dll.head
		for node.next != nil {
			result += "{" + strconv.Itoa((*node).value) + "} <-> "
			node = node.next
		}
		result += "{" + strconv.Itoa((*node).value) + "}"
		return result
	}
	return "nil"
}

func (dll *DoubleLinkedList) print() {
	fmt.Print(dll.toString())
}

// Retrieves a node from the double linked list at the given index
func (dll *DoubleLinkedList) get( index int) *listNode {
	if index >= dll.size {
		return nil
	}
	
	node := dll.head
	for index > 0 {
		if node == nil {
			break
		}
		node = node.next
		index--
	}
	return node
}

func (dll *DoubleLinkedList) insert(value int, index int) (err string) {
	if index < 0 || index > dll.size {
		return "Insertion index is out of bounds"
	}
	newNode := &listNode{
		value: value,
		next: nil,
		prev: nil,
	}
	
	if dll.head == nil {
		dll.head = newNode
		dll.tail = dll.head
	} else {
		node := dll.get(index)
		if index == 0 {
			newNode.next = node		// {new} -> {node}
			newNode.next.prev = newNode // {new} <-> {node}
			dll.head = newNode	// h* -> {new}
		} else if index == dll.size {
			dll.tail.next = newNode	// {t} -> {new}
			newNode.prev = dll.tail	// {t} <-> {new}
			dll.tail = newNode	// *t -> {new}
		} else {
			node.prev.next = newNode	// {prev} -> {new}
			newNode.prev = node.prev	// {prev} <-> {new}
			newNode.next = node // {prev} <-> {new} -> {node}
			node.prev = newNode	// {prev} <-> {new} <-> {node}
		}
	}
	dll.size++
	return ""
}

func (dll * DoubleLinkedList) remove(index int) string {
	if index < 0 || index > dll.size - 1 {
		return "Insertion index is out of bounds"
	}

	// Incase we only have 1 item left
	if (dll.size == 1) {
		dll.size = 0
		dll.head = nil
		dll.tail = nil
		return ""
	}
	
	toRemove := dll.get(index)
	
	// Removing from the ends
	if index == dll.size - 1 {
		toRemove.prev.next = nil	// {prev} -> nil
		dll.tail = toRemove.prev	// t* -> {prev}
	} else if index == 0 {
		toRemove.next.prev = nil	// nil <- {next}
		dll.head = toRemove.next	// h* -> {next}
	} else {
		toRemove.prev.next = toRemove.next	// {prev} -> {ndex}
		toRemove.next.prev = toRemove.prev	// {prev} <-> {next}
	}
	
	dll.size--
	return ""
}

// CreateDoubleLinkedList creates a double linked list
func CreateDoubleLinkedList() (DoubleLinkedList) {
	return DoubleLinkedList{ 
		size: 0,
		head: nil,
		tail: nil,
	}
}