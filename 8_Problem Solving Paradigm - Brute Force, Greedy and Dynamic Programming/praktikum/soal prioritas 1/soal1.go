package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(representBiner(2)) //[0,1,10]
	fmt.Println(representBiner(4)) //[0,1,10,11,100]
}

func representBiner(n int) [][]int {
	var res [][]int
	for i := 0; i <= n; i++ {
		binary := decimalToBiner(i)
		res = append(res, binary)
	}
	return res
}

func decimalToBiner(n int) []int {
	var res []int
	if n == 0 {
		return []int{0}
	}
	for n > 0 {
		res = append(res, n%2)
		n = n / 2
	}
	slices.Reverse(res)
	return res
}
