package main

import (
	"fmt"
)

func main() {
	res1 := merge([][]int{
		{1, 2, 5, 4, 3},
		{1, 2, 7, 8},
	})
	fmt.Println(res1) // [1,2,5,4,3,7,8]

	res2 := merge([][]int{
		{1, 2, 5, 4, 5},
		{7, 9, 10, 5},
	})
	fmt.Println(res2) // [1, 2, 5, 4, 7, 9, 10]

	res3 := merge([][]int{
		{1, 4, 5},
		{7},
	})
	fmt.Println(res3) // [1, 4, 5, 7]
}

func merge(data [][]int) []int {
	//TODO: write your code here
	tempArr := []int{}
	resArr := []int{}
	//Menggabungkan  arr 2 dimensi menjadi 1 dimensi
	for _, v := range data {
		for i := 0; i < len(v); i++ {
			tempArr = append(tempArr, v[i])
		}
	}
	//Mengurutkan array dan mengganti nilai duplikat menjadi 0
	for i := 0; i < len(tempArr); i++ {
		for j := 0; j < len(tempArr)-1; j++ {
			if tempArr[j] > tempArr[j+1] {
				temp := tempArr[j]
				tempArr[j] = tempArr[j+1]
				tempArr[j+1] = temp
			}
			if tempArr[j] == tempArr[j+1] {
				tempArr[j+1] = 0
			}
		}
	}
	//Menghapus nilai 0
	for _, num := range tempArr {
		if num != 0 {
			resArr = append(resArr, num)
		}
	}
	return resArr
}
