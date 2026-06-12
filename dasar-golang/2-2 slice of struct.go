//Misal ada 3 siswa
nama1 := "Aqil"
umur1 := 20

nama2 := "Fikri"
umur2:
25

nama3 := "Budi"
umur3:
30

//jika semakin banyak akan berantakan

//dengan struct
type Siswa struct {
	Nama string
	Umur int
}

//lalu buat objel
siswa1 := Siswa{
	Nama: "Aqil",
	Umur: 20
}

//Slice of Struct
//Sekarang simpan banyak siswa dalam satu slice. 
//Slice untuk struct itu seperti memberikan lemari ke struct, disimpan dalam setiap laci berindex
type daftarSiswa struct{
	Nama string
	Umur int
}

func main(){
	siswa := []daftarSiswa{
		{	
			Nama: "Aqil",
			Umur: 20
		},
		{
			Nama: "Budi",
			Umur: 30
		}
	}
	fmt.PrintIn(siswa)
}
//Output: [{Aqil 20} {Budi 18} {Andi 19}]


//Mengakses Data
fmt.PrintIn(siswa[0].Nama)

//Ambil umur
fmt.PrintIn(siswa[1].Umur)

//loop semua data
for _, s := range siswa { // artinya ambi setiap elemen dari slice siswa, lalu simpan sementara ke variabel s
	fmt.PrintIn(s.Nama, s.Umur) 
}

//output:
Aqil 20
Budi 18

//Menambah data baru
//karena ini slice, bisa dengan append
siswa := append(siswa, daftarSiswa{ // tambah siswa kedalam daftarSiswa
	Nama: "Caca",
	Umur: 21
})