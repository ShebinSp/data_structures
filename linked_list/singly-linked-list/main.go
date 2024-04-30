package main

import "fmt"

type Node struct {
	data any
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Insert(data interface{}) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (list *LinkedList) Display() {
	current := list.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

// Third arg is bool, true deletes all matches, false deletes first match.
func (list *LinkedList) Delete(data interface{}, deleteAll bool) {
	current := list.head
	if current == nil {
		fmt.Println("Linked List is empty")
	} else {
		for current != nil {
			if current.data == data {
				fmt.Printf("Type: %T\n",current.data)
				current.data = current.next.data
				current.next = current.next.next
				if !deleteAll {
					break 
				}

			} else {
				current = current.next
			}
		}
	}
}

// Third arg is bool, true updates all matches, false updates first match.
func (list *LinkedList) Update(data, newData interface{}, updateAll bool) {
	current := list.head
	if current == nil {
		fmt.Println("Linked List is empty")
	} else {
		for current != nil {
			if current.data == data {
				current.data = newData
				if !updateAll {
					break 
				}
			}
			current = current.next
		}

	}
}

func main() {
	myList := LinkedList{}
	myList.Insert("hai")
	myList.Insert(1)
	myList.Insert("a")
	myList.Insert(11.11)
	myList.Insert("a")
	myList.Insert(1)

	myList.Display()
	fmt.Println("----------------------------")
	myList.Delete(1,false)
	myList.Display()
	fmt.Println("----------------------------")
	myList.Update(1, 3, false)
	myList.Display()
}
