package main

import "fmt"

type Student struct {
	Name         string
	MathScore    int
	ScienceScore int
	EnglishScore int
}

func main() {
	students := []Student{
		{"jamie", 67, 88, 90},
		{"michael", 77, 77, 90},
		{"aziz", 100, 65, 88},
		{"jamal", 50, 80, 75},
		{"eric", 70, 80, 65},
		{"john", 80, 76, 68},
	}

	getInfo(students)
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
