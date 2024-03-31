# Resume Pertemuan 17 - Gorm

### Resume materi  17 Gorm

***Gorm** adalah salah satu ORM dari golang yang sanagat memudahkan kita untuk 

Untuk melakukan retrive/mendapatkan data dapat melalui 3 cara
1. URL Params
2. Query Params
3. Form Value

Dalam echo kita dimudahkan dalam melakukan binding data, kita hanya mengarahkan ke dalam sebuah struct dan memberikannya alamat memory struct tersebut
```
err = c.Bind(&food)
```
Binding mengembalikan sebuah error apabila dibutuhkan untuk handling error