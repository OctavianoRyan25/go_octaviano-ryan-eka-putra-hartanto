# Resume Pertemuan 16 - Echo

### Resume materi  16 Echo

***Echo** adalah salah satu framework yang kuat dan serbaguna untuk membangun aplikasi web yang skalabilitas dan kinerjanya tinggi dalam bahasa pemrograman Go. Ini mengikuti prinsip kesederhanaan, fleksibilitas, dan kinerja untuk menyediakan perangkat yang efisien bagi developer untuk membangun aplikasi web yang kuat.

Untuk melakukan retrive/mendapatkan data dapat melalui 3 cara
1. URL Params
2. Query Params
3. Form Value

Dalam echo kita dimudahkan dalam melakukan binding data, kita hanya mengarahkan ke dalam sebuah struct dan memberikannya alamat memory struct tersebut
```
err = c.Bind(&food)
```
Binding mengembalikan sebuah error apabila dibutuhkan untuk handling error