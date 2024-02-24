package main

import "fmt"

func main() {
	//
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 1; i <= 5; i++ {
		for j := 5; j >= i; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 1; i <= 5; i++ {
		for k :=4 ; k >= i; k-- {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
