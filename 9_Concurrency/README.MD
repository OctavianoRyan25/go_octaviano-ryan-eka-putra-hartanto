# Resume Pertemuan 9 - Concurrent Program

### Resume materi  9 Concurrent Program

***Concurrent program*** adalah kemampuan suatu program untuk menjalankan beberapa tugasnya secara bersamaan. didalam konkurensi dapat menerapkan paralelism. Jadi, tugas tugas tadi dapat dibagi lagi kedalam thread yang berbeda

***Goroutine*** adalah sebuah fungsi atau method yang dijalankan secara konkurensi dengan fungsi dan method lainnya
```
go func(){
    //Your Code

}()
```

***Channel*** adalah sebuah objek yang digunakan agar goroutine dapat berkomunikasi satu sama lain
```
c := make(chan string) //inisiasi channel
c <- "Lorem" //Mengisi channel dengan string "Lorem"
```