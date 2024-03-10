package main

import "fmt"

func main() {
	fmt.Println(fibX(5))  // 12
	fmt.Println(fibX(10)) // 143
}

func fibX(num int) (res int) {
	if num <= 1 {
		return num
	} else {
		return fibX(num-2) + fibX(num-1) + 1
	}
}
