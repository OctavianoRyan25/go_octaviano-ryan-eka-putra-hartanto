package main

import "fmt"

type Stack struct {
	item int
	next *Stack
	prev *Stack
}

var stack, current, top, bottom *Stack

func main() {
	print()
	push(1)
	print()
	push(2)
	print()
}

func isEmpty() bool {
	return top == nil
}

func push(item int) {
	current = &Stack{item, nil, nil}
	if isEmpty() {
		top = current
		bottom = current
	} else {
		current.next = top
		top = current
	}
}

func print() {
	if top == nil {
		fmt.Println("Stack is empty")
	}
	current = top
	for current != nil {
		println(current.item)
		current = current.next
	}
}

func pop(item int) {
	current = &Stack{item, nil, nil}
	if isEmpty() {
		top = current
		bottom = current
	} else {
		current.next = nil
		top = current
	}
}
