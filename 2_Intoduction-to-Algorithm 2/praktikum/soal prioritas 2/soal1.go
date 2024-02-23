package main

import "fmt"

func main() {
	var param int = 35
	for i := 1; i <= param; i++ {
		if param%i == 0 {
			if i%2 == 0 {
				fmt.Println("I")
			}else{
				fmt.Println("O")
			}
		}
	}
}
