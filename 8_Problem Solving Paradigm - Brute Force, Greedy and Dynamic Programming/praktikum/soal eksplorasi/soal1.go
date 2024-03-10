package main

import (
	"fmt"
)

func main() {
	fmt.Println(decimalToRoman(21)) // "XXI"
}

func decimalToRoman(n int) string {
	roman := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	var res string
	target := n
	for target > 0 {
		for k, v := range roman {
			if target >= k {
				target -= k
				res += v
				break
			}
		}
	}
	return res
}
