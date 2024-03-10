package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}

func Frog(jumps []int) int {
	n := len(jumps)

	dp := make([]int, n)

	dp[0] = 0
	dp[1] = int(math.Abs(float64(jumps[1] - jumps[0])))

	for i := 2; i < n; i++ {
		dp[i] = min(dp[i-1]+int(math.Abs(float64(jumps[i]-jumps[i-1]))), dp[i-2]+int(math.Abs(float64(jumps[i]-jumps[i-2]))))
	}

	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
