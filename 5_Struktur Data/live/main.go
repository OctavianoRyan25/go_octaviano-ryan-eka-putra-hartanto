package main

import "fmt"

func main() {
	// fmt.Println("Hello, World!")
	fmt.Println(concatenation([]int{1, 2, 3, 4, 5}))
}

func concatenation(data []int) []int{
	resSlice := []int{}
	resSlice = append(resSlice, data...)
	for _, v := range data {
		resSlice = append(resSlice, v)
	}
	return resSlice
}

