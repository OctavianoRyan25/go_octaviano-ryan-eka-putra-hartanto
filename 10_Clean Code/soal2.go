package main

import "fmt"

type HashSet struct {
	Data map[int]int
}

func (hashset HashSet) add(i int) {
	hashset.Data[i] = 1
}

func (hashset HashSet) get() []int {
	res := []int{}

	for key := range hashset.Data {
		res = append(res, key)
	}

	return res
}

func (hashset HashSet) remove(key int) {
	delete(hashset.Data, key)
}

func main() {
	scoreData := HashSet{
		Data: make(map[int]int),
	}

	scoreData.add(1)
	scoreData.add(2)
	scoreData.add(1)
	scoreData.add(3)
	scoreData.add(4)
	scoreData.add(5)

	scoreData.remove(100)

	fmt.Println(scoreData.get())
}
