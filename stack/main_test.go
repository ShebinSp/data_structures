package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestStack_Push(t *testing.T) {
	myStack := &Stack{}

	// Push 5 items to the stack
	for i := 1; i <= 5; i++ {
		err := myStack.Push(i)
		if err != nil {
			t.Errorf("Pushing item %d should not return an error, got %v", i, err)
		}
	}

	// Attempt to push one more item, which should return an error
	err := myStack.Push(6)
	cstmErr := "Stack overflow"

	cstmErr = fmt.Sprintf("%s: Stack overflow on line %d", "an error occurred on Push", 21)

	expectedErr := errors.New(cstmErr)
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("Pushing item 6 should return error: %v, got: %v", expectedErr, err)
	}
}

func TestStack_Pop(t *testing.T) {
	myStack := &Stack{}

	// Push an item to the stack
	myStack.Push(1)

	// Pop the item from the stack
	item, err := myStack.Pop()
	if err != nil {
		t.Errorf("Should not return an error, got: %v", err)
	}
	if item != 1 {
		t.Errorf("Popped item should be 1, got: %v", item)
	}

	// Pop from an empty stack, which should return an error
	_, err = myStack.Pop()
	expectedErr := errors.New("Stack is empty")
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("Popping from empty stack should return error: %v, got: %v", expectedErr, err)
	}
}

func TestStack_Peek(t *testing.T) {
	myStack := &Stack{}

	// Peek from an empty stack, which should result in an error
	_, err := myStack.Peek()
	expectedErr := errors.New("an error occurred on Peek : stack is empty")

	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("Peeking from empty stack should return error: %v, got: %v", expectedErr, err)
	}

	// Push an item to the stack
	myStack.Push(1)

	// Peek the item from the stack
	item, err := myStack.Peek()
	if err != nil {
		t.Errorf("Peeking should not return an error, got: %v", err)
	}
	if item != 1 {
		t.Errorf("Peeked item should be 1, got: %v", item)
	}
}

func TestStack_PeekOnIndex(t *testing.T) {
	myStack := &Stack{}

	// PeekOnIndex from an empty stacks, which should result in an error
	_, err := myStack.PeekOnIndex(0)

	expectedErr := errors.New("an error occurred on Peek : stack is empty")

	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("PeekOnIndex from empty stack should retrun error: %v, got: %v", expectedErr, err)
	}

	// Push items to the stack
	for i := 1; i <= 5; i++ {
		myStack.Push(i)
	}

	// PeekOnIndex with valid index
	item, err := myStack.PeekOnIndex(1)
	if err != nil {
		t.Errorf("PeekonIndex should not return an error, got: %v", err)
	}
	if item != 2 {
		t.Errorf("Peeked item at index 1 should be 2, got %v", item)
	}

	// PeekOnIndex with invalid index
	_, err = myStack.PeekOnIndex(5)
	expectedErr = errors.New("an error occurred on Peek : invalid index")
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("PeekOnIndex with invalid index should return error: %v, got: %v", expectedErr, err)
	}
}
