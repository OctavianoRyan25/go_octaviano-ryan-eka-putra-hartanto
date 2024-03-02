package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(pickVocals("sepulsa mantap jiwa")) // ltrrcdmy
}

func pickVocals(sentence string) string {
	//TODO: your code here
	vocal := "aiueo "
	res := []string{}
	for _, value := range sentence {
		// Apabila mengandung huruf vokal maka masukkan ke dalam res
		if strings.Contains(vocal, string(value)) {
			res = append(res, string(value))
		}
	}
	return strings.Join(res, "")
}
