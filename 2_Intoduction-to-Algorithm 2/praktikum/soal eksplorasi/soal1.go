package main

import "fmt"

func main() {
	var param1 string = "pulsa"
	var param2 string = "palsu"
	if isAnagram(param1, param2) {
		fmt.Println("Anagram")
	}else{
		fmt.Println("Bukan Anagram")
	}
}

func isAnagram(param1 string, param2 string) bool{
	if len(param1) != len(param2) {
		return false
	}
	m := make(map[rune]int)
	for _, r := range param1 {
		m[r]++
	}
	for _, r := range param2 {
		m[r]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}