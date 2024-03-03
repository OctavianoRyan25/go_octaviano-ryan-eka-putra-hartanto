package main

import (
	"fmt"
)

type Item struct {
	item   int
	lanjut *Item
}

type Stack struct {
	top *Item
}

func (s *Stack) push(item int) {
	// Cek apakah item sudah ada di stack
	if !s.isDuplicate(item) {
		mhsBaru := &Item{
			item:   item,
			lanjut: s.top,
		}
		s.top = mhsBaru
		fmt.Printf("Item %v berhasil ditambahkan\n", item)
	} else {
		fmt.Println("Item sudah ada di stack")
	}
}

func (s *Stack) pop() (*Item, string) {
	res := "Item berhasil dihapus"
	if s.top == nil {
		return nil, res
	}
	popped := s.top
	s.top = s.top.lanjut
	return popped, res
}

func (s *Stack) values() {
	fmt.Println("Print items di stack :")
	current := s.top
	for current != nil {
		fmt.Print(current.item, " ")
		current = current.lanjut
	}
}

func (s *Stack) isEmpty() bool {
	return s.top == nil
}

func (s *Stack) isDuplicate(item int) bool {
	current := s.top
	for current != nil {
		if current.item == item {
			return true
		}
		current = current.lanjut
	}
	return false
}

func main() {
	// Inisiasi	stack
	stack := &Stack{}

	//Push item ke dalam stack
	stack.push(10)
	stack.push(20)
	stack.push(22)
	stack.push(20)
	stack.push(10)

	//Print stack
	stack.values()

	fmt.Println()

	// Pop top stack
	popped1, res := stack.pop()
	if popped1 != nil {
		fmt.Println(res)
	}
}
