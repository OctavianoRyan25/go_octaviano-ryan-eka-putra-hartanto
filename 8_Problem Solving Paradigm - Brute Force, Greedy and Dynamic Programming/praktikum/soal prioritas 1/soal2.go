package main

import "fmt"

func main() {
	fmt.Println(numRows(5)) // [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
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
