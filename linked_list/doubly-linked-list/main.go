package main

import "fmt"

type Node struct {
	data any
	prev *Node
	next *Node
}

type doublyLinkedList struct {
	head *Node
}

// Inset data into the last
func (l *doublyLinkedList) InsertLast(data any) {
	newNode := &Node{data: data}

	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head

		for current.next != nil {
			current = current.next
		}
		current.next = newNode
		newNode.prev = current
	}
}

// Insert data in to the front of the list
func (l *doublyLinkedList) InsertFront(data any) {
	newNode := &Node{data: data}
	if l.head == nil {
		l.head = newNode
	} else {
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
	}

}

// Traverse the list from front
func (l *doublyLinkedList) TraverseFront() {
	current := l.head

	for current != nil {
		fmt.Printf("%v \t", current.data)
		current = current.next
	}

	fmt.Println()
}

// Traverse the list from back
func (l *doublyLinkedList) TraverseBack() {

	if l.head == nil {
		fmt.Println("List is empty")
	} else {
		current := l.head

		for current.next != nil {
			current = current.next
		}

		for current != nil {
			fmt.Printf("%v\t", current.data)
			current = current.prev
		}
	}

	fmt.Println()
}

// Number of nodes in the list
func (l *doublyLinkedList) TotalNodes() int {
	count := 0
	current := l.head
	
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func (l *doublyLinkedList) Delete(data any) {
	if l.head == nil {
		fmt.Println("The list is empty")
	}

	if l.head.data == data {
		l.head = l.head.next
		if l.head != nil {
			l.head.prev = nil
		}
		return
	}

	current := l.head

	for current.next != nil {
		if current.next.data == data {
			nodeToDel := current.next
			current.next = nodeToDel.next
			if nodeToDel.next != nil {
				nodeToDel.next.prev = current
			}
			return
		}
		current = current.next
	}
	fmt.Println("Node not found")
}
func main() {
	newDLL := doublyLinkedList{}
	newDLL.InsertFront(1)
	newDLL.InsertFront(2)
	newDLL.InsertFront(3)
	newDLL.InsertFront(4)
	newDLL.TraverseFront()
	newDLL.TraverseBack()
	newDLL.InsertLast(7)
	newDLL.InsertFront(3)
	newDLL.TraverseBack()
	newDLL.TraverseFront()
	newDLL.TraverseBack()
	fmt.Println("No. of items in the list: ", newDLL.TotalNodes())
	newDLL.TraverseFront()
	newDLL.Delete(2)
	newDLL.TraverseFront()
	fmt.Println("No. of items in the list: ", newDLL.TotalNodes())
}
