package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getItem([]int{10, 10, 5, 30, 40, 10, 5}, 80)) // [40 30 10]
	fmt.Println(getItem([]int{50, 20, 10, 25, 25}, 100))      // [50 25 25]
}

func getItem(items []int, target int) []int {
	sort.Slice(items, func(i, j int) bool {
		return items[i] > items[j]
	})
	res := []int{}
	batas := 0
	for target != 0 {
		found := false
		for i := batas; i < len(items); i++ {
			if target >= items[i] {
				target -= items[i]
				res = append(res, items[i])
				found = true
				break
			}
		}
		if !found {
			break
		}
		batas++
	}
	return res
}
