package main

import "fmt"

func main() {
	calculateAll := calculateBudget(51)+calculateDay(10)+calculateDificulty(3)
	if calculateAll <= 2 {
		fmt.Println("impposible")
	}else if calculateAll >= 3 && calculateAll <= 9{
		fmt.Println("low")
	}else if calculateAll >= 10 && calculateAll <= 16{
		fmt.Println("medium")
	}else if calculateAll >= 17 && calculateAll <= 24{
		fmt.Println("high")
	}
}

func calculateBudget(budget int) (score int){
	if budget <= 50 {
		score = 4
	}else if budget >= 51 && budget <= 80{
		score = 3
	}else if budget >= 81 && budget <= 100{
		score = 2
	}else{
		score = 1
	}
	return score
}

func calculateDay(budget int) (score int){
	if budget <= 20 {
		score = 10
	}else if budget >= 21 && budget <= 30{
		score = 5
	}else if budget >= 31 && budget <= 50{
		score = 2
	}else{
		score = 1
	}
	return score
}

func calculateDificulty(budget int) (score int){
	if budget <= 3 {
		score = 10
	}else if budget >= 4 && budget <= 6{
		score = 5
	}else if budget >= 8 && budget <= 10{
		score = 1
	}else{
		score = 0
	}
	return score
}

