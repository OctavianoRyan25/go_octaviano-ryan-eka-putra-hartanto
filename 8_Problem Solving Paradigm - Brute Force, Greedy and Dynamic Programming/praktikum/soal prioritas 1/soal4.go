package main

import "fmt"

func main() {
	SimpleEquations(6, 6, 14)
}

func SimpleEquations(a, b, c int) {
	// your code here
	var x, y, z int
	for x = 0; x <= 100; x++ {
		for y = 0; y <= 100; y++ {
			for z = 0; z <= 100; z++ {
				if x+y+z == a && x*y*z == b && x*x+y*y+z*z == c {
					fmt.Println(x, y, z)
					return
				}
			}
		}
	}
}
