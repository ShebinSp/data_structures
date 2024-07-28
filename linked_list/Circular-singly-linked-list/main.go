package main

import "fmt"

type Node struct {
	data any
	next *Node
}

type CSLL struct {
	head *Node
}

func (list *CSLL) Insert(data any) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
		newNode.next = list.head
		return
	}

	current := list.head

	for current.next != list.head{
		current = current.next
	}
	current.next = newNode
	newNode.next = list.head
}

func (list *CSLL) Traverse() {
	if list.head == nil {
		fmt.Println("list is empty")
		return
	}

	current := list.head
	first := true

	for first || current != list.head {
		first = false
		fmt.Printf("%v\t", current.data)
		current = current.next
	}
	fmt.Println()
}

func main() {
	myList := &CSLL{}

	myList.Insert(1)
	myList.Insert(2)
	myList.Insert(3)
	myList.Insert(4)
	myList.Traverse()
}
