package main

import "fmt"

func main() {
	fmt.Println(primeSum([]int{1, 2, 4, 5, 12, 19, 30})) // 26
	fmt.Println(primeSum([]int{45, 17, 44, 67, 11}))     // 95
	fmt.Println(primeSum([]int{15, 12, 9, 10}))          // 0
}

func primeSum(numbers []int) int {
	res := 0
	//Looping untuk mengecek apakah angka tersebut adalah bilangan prima dan menambahkannya ke res
	for _, v := range numbers {
		if primeNumber(v) {
			res += v
		}
	}
	return res
}

func primeNumber(number int) bool {
	if number <= 1 {
		return false
	}
	if number == 2 {
		return true
	}
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
