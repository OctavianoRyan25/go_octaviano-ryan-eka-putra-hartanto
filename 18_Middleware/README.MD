# Resume Pertemuan 17 - Middleware

### Resume materi  17 Middleware

***Gorm** adalah sebuah lapisan perangkat lunak yang berfungsi sebagai jembatan antara beberapa aplikasi atau sistem. Dalam konteks pengembangan perangkat lunak, middleware berada di antara aplikasi klien dan server. Middleware memiliki peran penting dalam memfasilitasi komunikasi antara kedua aplikasi tersebut dengan menyediakan layanan seperti autentikasi, otorisasi, caching, logging, dan lain sebagainya.

Dalam Echo dibagi menjadi 2 tipe
1. Basic auth : CORS, JWT, CSRF
2. Key Auth : Session, Static, Logger

Logger menampilkan informasi HTTP request
JWT terdiri dari 3 bagian
1. HEADER yang berisi algoritma
2. PAYLOAD yang berisi data pengguna(usahakan tidak menggunakan data yang krusial seperti password)
3. VERIFY SIGNATURE yang berisi hasil dari hash atau gabungan dari isi encode Header dan Payload lalu ditambahkan dengan kode rahasia (secret). Signature ini berguna untuk memverifikasi bahwa header maupun payload yang ada dalam token tidak berubah dari nilai aslinya