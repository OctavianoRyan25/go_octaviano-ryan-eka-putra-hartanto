package main

import "fmt"

func main() {
	fmt.Println(updateTimes([]int32{1, 2, 3, 4, 5, 9, 91, 6}, []int32{5, 2, 3, 4, 1, 9, 91, 6})) // 5
}

func updateTimes(signalOne []int32, signalTwo []int32) int32 {
	// TODO: Write your code here
	//countArr to store the same number
	countArr := []int32{}
	count := int32(0)
	for i := 0; i < len(signalOne); i++ {
		if signalOne[i] == signalTwo[i] {
			countArr = append(countArr, signalOne[i])
		}
	}
	//count the number of increasing numbers
	if len(countArr) > 0 {
		count = 1
		for i := 0; i < len(countArr)-1; i++ {
			if countArr[i] < countArr[i+1] {
				count++
			}
		}
	}
	return count
}
