package main

import "fmt"

var data map[int]int

func main() {

	data = make(map[int]int)

	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144
}

func fibonacci(number int) int {
	if number <= 1 {
		return number
	}

	// Cek apakah sudah menghitung angka Fibonacci ini sebelumnya
	if v, ok := data[number]; ok {
		return v
	}

	// Hitung angka Fibonacci
	data[number] = fibonacci(number-1) + fibonacci(number-2)
	return data[number]
}
