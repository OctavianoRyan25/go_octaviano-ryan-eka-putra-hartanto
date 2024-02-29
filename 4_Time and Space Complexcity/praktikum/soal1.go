package main

import "fmt"

func main() {
// 	fmt.Println(primeNumber(1000000007)) // true
//    fmt.Println(primeNumber(13))         // true
//    fmt.Println(primeNumber(17))         // true
//    fmt.Println(primeNumber(20))         // false
//    fmt.Println(primeNumber(35))         // false
	// fmt.Println(pow(2, 3))  // 8
	// fmt.Println(pow(5, 3))  // 125
	// fmt.Println(pow(10, 2)) // 100
	fmt.Println(pow(2, 5))  // 32
	// fmt.Println(pow(7, 3))  // 343
	// fmt.Println(pow2(2, 5))  // 32
}

func primeNumber(number int) bool {
	if number <= 2{
		return true
	}
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func pow(x, n int) int {
	var temp int = 1
	for i := 0; i < n; i++ {
		temp *= x
	}
	return temp
}

func pow2(x, n int) int {
	var temp int = 1
	if n == 1 {
		return temp*x
	}
	if n >= 2{
		temp = x*pow2(x, n-1)
	}
	return temp
}