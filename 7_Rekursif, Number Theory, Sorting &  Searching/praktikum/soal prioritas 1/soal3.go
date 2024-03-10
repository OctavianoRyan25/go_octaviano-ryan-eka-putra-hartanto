package main

import "fmt"

func main() {
	fmt.Println(getSequence(4))   // 10
	fmt.Println(getSequence(15))  // 120
	fmt.Println(getSequence(100)) // 5050
}

func getSequence(n int) int {
	if n <= 1 {
		return n
	} else {
		return n + getSequence(n-1)
	}
}
