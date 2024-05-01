package main

import "fmt"

type Queue struct {
	items []interface{}
}

// Enqueue adds an element to the end of the queue
func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the first element from the queue
func (q *Queue) Dequeue() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// Front returns the first element of the queue without removing it
func (q *Queue) Front() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	return q.items[0]
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of elements in the queue
func (q *Queue) Size() int {
	return len(q.items)
}

func main() {

	q := Queue{}

	q.Enqueue("s")
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue("a")
	q.Enqueue("b")
	q.Enqueue("c")

	fmt.Println("Queue size:", q.Size())
	fmt.Println("Front element:", q.Front())

	fmt.Println("Dequeuing elements:")
	for !q.IsEmpty() {
		fmt.Println(q.Dequeue())
	}

}
