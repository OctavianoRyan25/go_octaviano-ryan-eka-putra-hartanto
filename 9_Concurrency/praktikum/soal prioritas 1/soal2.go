package main

func isPrime(primes chan int) {
	for i := 2; i <= 100; i++ {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes <- i
		}
	}
	close(primes)
}

func powPrimes(primes chan int, results chan int) {
	for prime := range primes {
		power := prime * prime
		results <- power
	}
	close(results)
}
func main() {
	primes := make(chan int)
	results := make(chan int)

	go isPrime(primes)
	go powPrimes(primes, results)

	for result := range results {
		println(result)
	}
}
