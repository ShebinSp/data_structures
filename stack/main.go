package main

import (
	"errors"
	"fmt"
	"runtime"
)

const (
	stackLimit        = 5
	ErrorPeekFailed = "an error occurred on Peek"
	ErrorPushFailed = "an error occurred on Push"
)

type Stack struct {
	items []interface{}
	index int
}

// Push add an item to the top of the stack or returns an error if occurred
func (s *Stack) Push(item interface{}) error {
	if s.index >= stackLimit { // stack is limited to hold only 5 elements inside
		cstmErr := "Stack overflow"
		_, _, line, ok := runtime.Caller(1)
		if ok {
			cstmErr = fmt.Sprintf("%s: Stack overflow on line %d", ErrorPushFailed, line)
		}
		return errors.New(cstmErr)
	}
	s.index++ // when each item added to the stack the index is incremented by 1
	s.items = append(s.items, item)
	return nil
}

// Pop removes the item in the top of the stack and returns the item removed or an error if occurred
func (s *Stack) Pop() (interface{}, error) {
	if s.index == 0 {
		return nil, errors.New("Stack is empty")
	}
	item := s.items[s.index-1]
	s.index--
	s.items = s.items[:s.index]
	return item, nil
}

// Peek returns the element in the top of the stack or an error if occurred
func (s *Stack) Peek() (interface{}, error) {
	if s.index == 0 {
		cstmErr := fmt.Sprintf("%s : stack is empty", ErrorPeekFailed)
		return nil, errors.New(cstmErr)
	}
	return s.items[s.index-1], nil
}

// PeelOnIndex returns an element at the passed index or an error if occurred
func (s *Stack) PeekOnIndex(idx int) (interface{}, error) {
	// checking the stack is empty or not
	if s.index == 0 {
		cstmErr := fmt.Sprintf("%s : stack is empty", ErrorPeekFailed)
		return nil, errors.New(cstmErr)
	} else if idx >= s.index {
		cstmErr := fmt.Sprintf("%s : invalid index", ErrorPeekFailed)
		return nil, errors.New(cstmErr)
	}
	return s.items[idx], nil
}

func main() {

	myStack := &Stack{}

	// Peek() an error will be returned
	if result, err := myStack.Peek(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// if err := myStack.Push(1); err != nil {
	// 	fmt.Println(err)
	// }
	// if err := myStack.Push(2); err != nil {
	// 	fmt.Println(err)
	// }
	// if err := myStack.Push(1.12); err != nil {
	// 	fmt.Println(err)
	// }
	// if err := myStack.Push("car"); err != nil {
	// 	fmt.Println(err)
	// }
	// if err := myStack.Push("bike"); err != nil {
	// 	fmt.Println(err)
	// }

	for i := 1; i <= 6; i++ {
		err := myStack.Push(i)
		fmt.Println("err: ", err)
	}
	if err := myStack.Push("bike"); err != nil {
		fmt.Println(err)
	}
	if result, err := myStack.Peek(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	if result, err := myStack.PeekOnIndex(2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The element: ", result)
	}

	if result, err := myStack.Pop(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result, " removed from stack")
	}

	if result, err := myStack.PeekOnIndex(6); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	if result, err := myStack.Peek(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The element in top of the stack: ", result)
	}
}
