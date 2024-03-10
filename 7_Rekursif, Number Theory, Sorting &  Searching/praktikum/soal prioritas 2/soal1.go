package main

import (
	"fmt"
)

func main() {
	fmt.Println(catalan(7))  // 429
	fmt.Println(catalan(10)) // 16796
}

func catalan(n int) int {
	if n <= 1 {
		return 1
	}
	res := 0
	for i := 0; i < n; i++ {
		res += catalan(i) * catalan(n-i-1)
	}
	return res
}
