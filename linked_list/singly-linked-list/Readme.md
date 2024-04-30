# Go Linked List Implementation
**This repository contains an implementation of a singly linked list data structure in Go.**

## Table of Contents
* Overview
* Usage
* Methods
* Contributing

## Overview
This Go program provides an implementation of a singly linked list, a fundamental data structure in computer science. It includes methods for insertion, deletion, updating, and displaying elements in the linked list.

## Usage

* To use this linked list implementation, follow these steps:

1. Clone the repository:

`git clone https://github.com/ShebinSp/go-linked-list.git`

2. Navigate to the project directory:

`cd go-linked-list`

3. Run the main Go program:

`go run main.go`


## Methods
### Insert
* Inserts a new element at the end of the linked list.

`func (list *LinkedList) Insert(data any)`

### Delete
* Deletes the first occurrence of a specified element from the linked list.

`func (list *LinkedList) Delete(data any, deleteAll bool)`

### Update

* Updates the first occurrence or all occurrences of a specified element in the linked list with a new value.

`func (list *LinkedList) Update(data, newData any, updateAll bool)`


### Display
* Displays the elements of the linked list.

`func (list *LinkedList) Display()`


## Contributing
**Contributions to this project are welcome! Feel free to open issues or pull requests with any improvements or bug fixes.**

