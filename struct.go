package main

import "fmt"

type Mahasiswa struct {
	Nama string
	Umur int
}

func (m Mahasiswa) Perkenalan() {
	fmt.Printf("Halo, saya %s dan umur saya %d tahun.\n", m.Nama, m.Umur)
}

func main() {
	aqil := Mahasiswa{Nama: "Aqil", Umur: 20}
	aqil.Perkenalan()
}
