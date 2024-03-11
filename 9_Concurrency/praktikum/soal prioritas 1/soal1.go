package main

import (
	"fmt"
	"time"
)

func reverse(s string) {
	var result string
	go func() {
		for i := 0; i < len(s); i++ {
			for j := len(s) - 1; j >= i; j-- {
				result += string(s[j])
				fmt.Println(result)
				time.Sleep(3 * time.Second)
			}
		}
	}()
	time.Sleep(15 * time.Second)
}

func main() {
	reverse("rusak") //kasur
}
