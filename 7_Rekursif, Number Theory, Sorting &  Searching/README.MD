# Resume Pertemuan 7 - Recursive, Number Theory, Sorting, Searching

### Resume materi Recursive, Number Theory, Sorting, Searching

**Recursive** adalah sebuah fungsi yang memanggil dirinya sendiri
```
func faktorial(n int){
    if n==1{
        return n
    }else{
        return n*faktorial(n-1)
    }
}
```

**Number Theory** adalah cabang dari ilmu matematika yang mempelajari mengenai angka contohnya pencarian bilangan prima, FPB, KPK, Faktorisasi prima
```
func primeNumber(number int) bool {
	if number <= 1 {
		return false
	}
	if number <= 3 {
		return true
	}
	if number%2 == 0 || number%3 == 0 {
		return false
	}
	for i := 5; i*i <= number; i += 6 {
		if number%i == 0 || number%(i+2) == 0 {
			return false
		}
	}
	return true
}
```

**SEARCHING** adalah proses mencari sebuah nilai berdasarkan list dari nilai yang diberi. Pencarian ini banyak jenisnya contohnya linier search, binary search

**SORTING** adalah proses pengurutas suatu list angka dari yang terkecil ke terbesar maupun sebaliknya, contohnya yaitu bubble sort, counting sort


