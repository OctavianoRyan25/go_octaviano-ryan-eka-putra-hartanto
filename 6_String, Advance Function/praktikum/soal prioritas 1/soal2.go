package main

import "fmt"

func main() {
	fmt.Println(spinString("sepulsa"))
}

func spinString(sentence string) string {
	res := []byte{}
	for i := 0; i < len(sentence)/2; i++ {
		// masukkan huruf pertama dan terakhir ke dalam res
		res = append(res, sentence[i])
		res = append(res, sentence[len(sentence)-1-i])
	}
	// apabila panjang sentence ganjil maka masukkan huruf tengah ke dalam res index terkahir
	if len(sentence)%2 != 0 {
		res = append(res, sentence[len(sentence)/2])
	}
	return string(res)
}
