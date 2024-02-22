package main

import "fmt"

func main() {
	var str1 string = "pulsa"
	var str2 string = "palsu"
	if isAnagram(str1, str2) {
		fmt.Println("Anagram")
	}else{
		fmt.Println("Bukan Anagram")
	}
}

func isAnagram(str1 string, str2 string) bool{
	if len(str1) != len(str2) {
		return false
	}
	m := make(map[rune]int)
	for _, r := range str1 {
		m[r]++
	}
	for _, r := range str2 {
		m[r]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}