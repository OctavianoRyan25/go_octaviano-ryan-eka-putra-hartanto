package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	//Soal Prioritas 1
	var phi float64 = 3.14
	var r float64 = 7
	var t float64 = 10
	var volume float64 = phi * r * r * t
	fmt.Println("Volume Tabung adalah", volume)

	var nilai int = 10
	if (nilai >= 85 && nilai <= 100) {
		fmt.Println("A")
	}else if (nilai >= 70 && nilai <= 84) {
		fmt.Println("B")
	}else if (nilai >= 55 && nilai <= 69) {
		fmt.Println("C")
	}else if (nilai >= 40 && nilai <= 50) {
		fmt.Println("D")
	}else if (nilai >= 0 && nilai <= 39) {
		fmt.Println("D")
	}else {
		fmt.Println("Nilai invalid")
	}
	fmt.Println("Nilai adalah", nilai)

	for i := 1; i <= 100; i++ {
		if i%4==0 {
			fmt.Println(i*i)
		}else if i%7==0 {
			fmt.Println(i*i*i)
		}else if i%4==0 && i%7==0 {
			fmt.Println("buzz")
		}else{
			fmt.Println(i)
		}
	}

	//Soal Prioritas 2
	var angka int = 26
	for i := 1; i <= angka; i++ {
		if angka%i==0 {
			if i%2==0 {
				fmt.Println("I")
			}else{
				fmt.Println("O")
			}
		}
	}
}