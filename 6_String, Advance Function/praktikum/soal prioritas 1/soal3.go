package main

import (
	"fmt"
)

func main() {
	fmt.Println(groupPalindrome([]string{"kodok", "malam", "siang", "katak", "kodok"}))
}

func isPalindrome(param1 string) bool {
	original := param1
	reverse := ""

	// reverse string
	for i := len(original) - 1; i >= 0; i-- {
		reverse += string(original[i])
	}

	// apakah original sama dengan reverse
	if original != reverse {
		return false
	}
	return true
}

func groupPalindrome(words []string) []any {
	//TODO: your code here
	palindrome := []string{}
	resSlice := []any{}
	for i := 0; i < len(words); i++ {
		if isPalindrome(words[i]) {
			palindrome = append(palindrome, words[i])
		}
	}
	resSlice = append(resSlice, palindrome)
	for i := 0; i < len(words); i++ {
		if !isPalindrome(words[i]) {
			resSlice = append(resSlice, words[i])
		}
	}
	return resSlice
}
