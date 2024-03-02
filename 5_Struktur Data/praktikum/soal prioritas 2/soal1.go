package main

import "fmt"

func main() {
	fmt.Println(spinSlice([]int{1, 2, 3, 4, 5}))    // [1,5,2,4,3]
	fmt.Println(spinSlice([]int{6, 7, 8}))          // [6,8,7]
	fmt.Println(spinSlice([]int{8, 5, 4, 3, 2, 1})) // [8,1,5,2,4,3]
}

func spinSlice(numbers []int) []int {
	resArr := []int{}
	for i := 0; i < len(numbers)/2; i++ {
		resArr = append(resArr, numbers[i])
		resArr = append(resArr, numbers[len(numbers)-1-i])
	}
	if len(numbers)%2 != 0 {
		resArr = append(resArr, numbers[len(numbers)/2])

	}
	return resArr
}
