package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(representBiner(10))
	fmt.Println(numRows(5))
	SimpleEquations(6, 6, 14)
	fmt.Println(decimalToRoman(10))
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
	fmt.Println(res)
	return res
}

func numRows(n int) [][]int {
	resSlice := [][]int{}
	for i := 0; i < n; i++ {
		if i == 0 {
			slice := []int{1}
			resSlice = append(resSlice, slice)
		}
		if i == 1 {
			slice := []int{1, 1}
			resSlice = append(resSlice, slice)
		}
		if i > 1 {
			slice := []int{}
			for j := 0; j <= i; j++ {
				//ambil slice sebelumnya
				prevSlice := resSlice[i-1]
				//awal dan akhir selalu 1
				if j == 0 || j == i {
					slice = append(slice, 1)
				} else {
					slice = append(slice, prevSlice[j-1]+prevSlice[j])
				}
			}
			resSlice = append(resSlice, slice)
		}
	}
	return resSlice
}

func SimpleEquations(a, b, c int) {
	// your code here
	var x, y, z int
	for x = 0; x <= 100; x++ {
		for y = 0; y <= 100; y++ {
			for z = 0; z <= 100; z++ {
				if x+y+z == a && x*y*z == b && x*x+y*y+z*z == c {
					fmt.Println(x, y, z)
					return
				}
			}
		}
	}
}

func decimalToRoman(n int) string {
	//TODO : your code here
	return ""
}
