package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(primeFactors(20))   // 2, 2, 5
	fmt.Println(primeFactors(75))   // 3, 5, 5
	fmt.Println(primeFactors(1024)) // 2, 2, 2, 2, 2, 2, 2, 2, 2, 2
}

func primeFactors(n int) []int {
	res := []int{}

	// Khusus faktor prima 2
	for n%2 == 0 {
		res = append(res, 2)
		n = n / 2
	}

	// Faktor prima ganjil mulai dari 3 + 2.....
	for i := 3; i <= int(math.Sqrt(float64(n))); i = i + 2 {
		for n%i == 0 {
			res = append(res, i)
			n = n / i
		}
	}

	// Pastikan faktor terakhir dimasukkan ke dalam daftar faktor prima
	if n > 2 {
		res = append(res, n)
	}

	return res
}
