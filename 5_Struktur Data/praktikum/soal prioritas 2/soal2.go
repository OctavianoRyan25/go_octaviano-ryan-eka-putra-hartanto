package main

import "fmt"

func main() {
	fmt.Println(groupPrime([]int{2, 3, 4, 5, 6, 7, 8})) // [[2,3,5,7],4,6,8]
	fmt.Println(groupPrime([]int{15, 24, 3, 9, 5}))     // [[3,5],15,24,9]
	fmt.Println(groupPrime([]int{4, 8, 9, 12}))         // [4,8,9,12]
}

func groupPrime(numbers []int) []any {
	resArr := []any{}
	primeArr := []int{}
	// nonPrimeArr := []int{}

	for _, v := range numbers {
		if isPrime(v) {
			primeArr = append(primeArr, v)
		}
	}
	resArr = append(resArr, primeArr)
	for _, v := range numbers {
		if !isPrime(v) {
			resArr = append(resArr, v)
		}
	}
	// resArr = append(resArr, nonPrimeArr)
	return resArr
}

func isPrime(number int) bool {
	if number <= 2 {
		return true
	}
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
