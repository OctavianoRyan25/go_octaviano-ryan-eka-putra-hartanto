package main

import (
	"fmt"
	"math"
)

// func miniMaxSum(arr []int) {
// 	for i := 0; i <= len(arr)-1; i++ {
// 		arr[i] = 0
// 		for j := 0; j <= len(arr)-1; j++ {
// 			fmt.Print(arr[j])
// 		}
// 		fmt.Println()
// 	}
// 	return
// }

type Student struct {
	Name         string
	MathScore    int
	ScienceScore int
	EnglishScore int
}

func main() {
	// arr := []int{1, 2, 3, 4, 5}
	// miniMaxSum(arr)
	fmt.Println(fibonanchi(10))

	students := []Student{
		{"jamie", 67, 88, 90},
		{"michael", 77, 77, 90},
		{"aziz", 100, 65, 88},
		{"jamal", 50, 80, 75},
		{"eric", 70, 80, 65},
		{"john", 80, 76, 68},
	}
	getInfo(students)

	fmt.Println(getSequence(15))
	fmt.Println(catalan(7))
	fmt.Println(primeFactors(75))
	fmt.Println(getSequence2(25))
}

func fibonanchi(num int) (res int) {
	if num <= 1 {
		return num
	} else {
		return fibonanchi(num-2) + fibonanchi(num-1) + 1
	}
}

func getInfo(param []Student) {
	maxMath := 0
	maxScience := 0
	maxEnglish := 0
	maxAverrage := 0
	winnerMath := ""
	winnerScience := ""
	winnerEnglish := ""
	winnerAverrage := ""
	for _, v := range param {
		if v.MathScore > maxMath {
			maxMath = v.MathScore
			winnerMath = v.Name
		}
		if v.ScienceScore > maxScience {
			maxScience = v.ScienceScore
			winnerScience = v.Name
			//fmt.Println(v.Name, v.ScienceScore)
		}
		if v.EnglishScore > maxEnglish {
			maxEnglish = v.EnglishScore
			winnerEnglish = v.Name
			//fmt.Println(v.Name, v.EnglishScore)
		}
		averrage := (v.EnglishScore + v.MathScore + v.ScienceScore) / 3
		if averrage > maxAverrage {
			maxAverrage = averrage
			winnerAverrage = v.Name
		}
	}
	fmt.Println("Winner math", winnerMath, maxMath)
	fmt.Println("Winner math", winnerScience, maxScience)
	fmt.Println("Winner math", winnerEnglish, maxEnglish)
	fmt.Println("Winner averrage", winnerAverrage, maxAverrage)
}

func getSequence(n int) int {
	if n <= 1 {
		return n
	} else {
		return n + getSequence(n-1)
	}
}

func catalan(n int) int {
	if n <= 1 {
		return 1
	}
	res := 0
	for i := 0; i < n; i++ {
		res += catalan(i) * catalan(n-i-1)
	}
	return res
}

func primeFactors(n int) []int {

	res := []int{}
	if n%2 == 0 {
		res = append(res, 2)
		n = n / 2
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i = i + 2 {
		// While i divides n, print i and divide n
		for n%i == 0 {
			res = append(res, i)
			n = n / i
		}
	}
	if n > 2 {
		res = append(res, n)
	}
	return res
}

func getSequence2(n int) int {
	if n <= 1 {
		return n
	} else {
		return n + getSequence(n-1) + getSequence(n-1)
	}
}
