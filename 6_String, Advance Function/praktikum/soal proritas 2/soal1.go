package main

import "fmt"

type Pelanggan struct {
	nama   string
	alamat string
	sewa   []Kendaraan
}

type Kendaraan struct {
	merk  string
	tipe  string
	harga int
}

func (p *Pelanggan) dataPelanggan() {
	fmt.Println("Nama Pelanggan: ", p.nama)
	fmt.Println("Alamat Pelanggan: ", p.alamat)
}

func (p *Pelanggan) sewaKendaraan(k Kendaraan) {
	p.sewa = append(p.sewa, k)
	fmt.Println("Kendaraan berhasil disewa")
}

func (p *Pelanggan) kembalikanKendaraan(k Kendaraan) {
	param := k.tipe
	for i, v := range p.sewa {
		if v.tipe == param {
			p.sewa = append(p.sewa[:i], p.sewa[i+1:]...)
			fmt.Println("Kendaraan berhasil dikembalikan")
		}
	}
}

// func (p *Pelanggan) removeKendaraan(k Kendaraan) {
// 	for i, v := range p.sewa {
// 		if v.tipe == k.tipe {

// 	}
// }

func (p *Pelanggan) daftarKendaraan() {
	fmt.Println("Daftar Kendaraan yang disewa", p.nama)
	for _, k := range p.sewa {
		fmt.Println("Merk: ", k.tipe)
	}
}

func main() {
	k1 := Kendaraan{"Toyota", "Avanza", 200000}
	k2 := Kendaraan{"Daihatsu", "Xenia", 250000}

	p1 := Pelanggan{"Budi", "Jakarta", []Kendaraan{k1}}
	// fmt.Println(p1)
	p1.sewaKendaraan(k2)

	p1.daftarKendaraan()

	p1.kembalikanKendaraan(k1)

	p1.daftarKendaraan()
}
