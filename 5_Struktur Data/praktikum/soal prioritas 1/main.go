package main

import (
	"fmt"
)


func main() {
	data1 := []float64{1, 1, 2, 5, 6, 8, 12, 4, 5, 5, 5, 8, 9}

	fmt.Println("sum: ", sum(data1))       // 71.00
	fmt.Println("mean: ", mean(data1))     // 5.46
	fmt.Println("median: ", median(data1)) // 5.00
	fmt.Println("mode: ", mode(data1))     // 5.00

	data2 := []float64{6, 7, 1, 11, 8, 12, 6, 12, 1, 7, 8, 10, 14}

	fmt.Println("sum: ", sum(data2))       // 103.00
	fmt.Println("mean: ", mean(data2))     // 7.92
	fmt.Println("median: ", median(data2)) // 8.00
	fmt.Println("mode: ", mode(data2))     // 1.00
}

func merge(data [][]int) []int {
	//TODO: write your code here
	tempArr := []int{}
	resArr := []int{}
	//Menggabungkan  arr 2 dimensi menjadi 1 dimensi
	for _, v := range data {
		for i := 0; i < len(v); i++ {
			tempArr = append(tempArr, v[i])
		}
	}
	//Mengurutkan array dan mengganti nilai duplikat menjadi 0
	for i := 0; i < len(tempArr); i++ {
		for j := 0; j < len(tempArr)-1; j++ {
			if tempArr[j]>tempArr[j+1] {
				temp := tempArr[j]
				tempArr[j] = tempArr[j+1]
				tempArr[j+1] = temp
			}
			if tempArr[j]==tempArr[j+1] {
				tempArr[j+1] = 0
			}
		}
	}
	//Menghapus nilai 0
	for _, num := range tempArr {
        if num != 0 {
            resArr = append(resArr, num)
        }
    }
	return resArr
}

func primeSum(numbers []int) int {
	res := 0
	//Looping untuk mengecek apakah angka tersebut adalah bilangan prima dan menambahkannya ke res
	for _, v := range numbers {
		if primeNumber(v) {
			res += v
		}
	}
	return res
}

func primeNumber(number int) bool {
	if number < 2{
		return true
	}
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func sum(data []float64) float64 {
	//TODO: your code here
	res := 0.0
	for _, v := range data {
		res += v
	}
	return res
}

func mean(data []float64) float64 {
	//TODO: your code here
	res := sum(data)/float64(len(data))
	return res
}

func median(data []float64) float64 {
	//TODO: your code here
	arrangeArray(data)
	if len(data)%2 == 0 {
		return (data[len(data)/2]+data[len(data)/2+1])/2
	}
	return data[len(data)/2]
}

func mode(data []float64) float64 {
	frequency := make(map[int]int)
    // Menghitung frekuensi kemunculan tiap elemen dalam array
    for i := 0; i < len(data); i++ {
		frequency[int(data[i])]++
	}
	// Mencari modus
	var maxFrequency float64
    var modeKey float64
    for key, freq := range frequency {
        if float64(freq) > maxFrequency {
            maxFrequency = float64(freq)
            modeKey = float64(key) // Mengupdate nilai modus jika ditemukan frekuensi yang lebih besar
        }
    }
    return modeKey
}

func arrangeArray(data []float64) []float64 {
	resArr := []float64{}
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data)-1; j++ {
			if data[j]>data[j+1] {
				temp := data[j]
				data[j] = data[j+1]
				data[j+1] = temp
			}
		}
	}
	for _, num := range data {
        if num != 0 {
            resArr = append(resArr, num)
        }
    }
	return resArr
}