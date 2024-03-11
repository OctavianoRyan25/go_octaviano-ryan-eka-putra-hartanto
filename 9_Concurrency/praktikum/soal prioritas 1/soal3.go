package main

import "fmt"

func main() {
	composites := make(chan int)
	res := make(chan string)

	go isComposite(composites)
	go processComposites(composites, res)

	for result2 := range res {
		fmt.Println(result2)
	}

}

func isComposite(composites chan int) {
	for i := 2; i <= 100; i++ {
		isComposite := false
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isComposite = true
				break
			}
		}
		if isComposite {
			composites <- i
		}
	}
	close(composites)
}

// Function to calculate power of composite numbers and check if they are even or odd
func processComposites(composites chan int, results chan string) {
	for composite := range composites {
		pow := composite * composite
		if pow%2 == 0 {
			results <- "Ping"
		} else {
			results <- "Pong"
		}
	}
	close(results)
}
