package main

import (
	"fmt"
)

type Cat struct {
	Hewan
	Name string
	Age  int
}

type Hewan struct {
	Famili  string
	Genus   string
	Species int
}

func (h *Hewan) String() {
	fmt.Print("ini adalah car", h.Species)

}

func main() {
	// names := "alterra"

	// for _, name := range names {
	// 	fmt.Print(name, " ")
	// }

	//string to int
	// sample := "1a"

	// res, err := strconv.Atoi(sample)
	// if err != nil {
	// 	fmt.Println("Error")
	// }

	// fmt.Println(res)

	//int to string
	// number := 10
	// res := strconv.Itoa(number)
	// fmt.Println(res)

	//compare string
	// name1 := "alterra"
	// res := strings.ToUpper(name1)
	// fmt.Println(res)

	//split string
	// name := "alterra academy"
	// res := strings.Split(name, "a") //Menghikangkan huruf a
	// fmt.Println(res)

	//spointer
	number := "10"
	// pointer := &number
	number = "20"
	address := &number
	pointNumber := *address
	var pointer *string = &number

	fmt.Println(pointer)
	fmt.Println(*pointer)
	fmt.Println(address)
	fmt.Println(pointNumber)
	// cat := Cat{
	// 	Hewan: Hewan{
	// 		Famili:  "Felidae",
	// 		Genus:   "Felis",
	// 		Species: 1,
	// 	},
	// 	Name: "Kucing",
	// 	Age:  2,
	// }

	// fmt.Println(cat)
	// cat.String()
}
