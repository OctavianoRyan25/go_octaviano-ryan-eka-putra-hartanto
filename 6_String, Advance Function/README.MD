# Resume Pertemuan 6 - String, Advance Function , Pointer , Method, Struct and Interface

### Resume materi String, Advance Function , Pointer , Method, Struct and Interface

**STRING** adalah tipe data yang berisi suatu kumpulan karakter

**POINTER** adalah variabel yang menyimpan alamat memory dari variabel lain
```
name := "Lorem"
nameAddress := &name //0xc000024070(alamat memory variabel tersebut)
pointername := *nameAddress // "Lorem"
```

* operator = Mendeklarasikan pointer variabel
& operator = mengembalikan alamat dari variabel

**STRUCT** adalah sekumpulan tipe data yang berbeda dan mengandung beberapa variabel/fungsi
```
type struct_name struct{
    variable <tipe_data>
    ...
    variable <tipe_data>
}
```

**INTERFACE** adalah sekumpulan method yang hanya dapat digunakan oleh object
```
type interface_name interface{
    method_name <return>
    ...
    variable <tipe_data>
}
```
