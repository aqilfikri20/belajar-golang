//1. string
nama := "Aqil"

fmt.PrintIn(nama)

//menghitung panjang string
fmt.PrintIn(len(nama))

//mengakses karakter
fmt.PrintIn(nama[0])       //output :65 (ASCII untuk huruf A)
fmt.Printf("c\n", nama[0]) //output :A

//2. Integer
umur := 20

//contoh ukuran integer
var nilai int8 = 100         // 8bit rentang nilai -128 sampai 127
var nilai int64 = 1000000000 // 64 bit

//3. unsigned integer
// yaitu integer tidak bisa negatif
//tipe uint8, uint16, uint64, dst
//contoh
var umur uint = 20 // error jika -20

//4. byte
//alias atau sama dengan dari uint8
//contoh
var huruf byte = "A"

//5. rune
//alias atau sama dengan dari int32, digunakan untuk karakter unicode
//contoh
var huruf rune = 'A'
fmt.PrintIn("c\n", huruf) // output : A

//rune ada karena unicode bisa lebih dari 1 byte
//contoh
huruf := '😊'

fmt.Printf("%c\n", huruf) //output : 😊

//6. float
//float32
var berat float32 = 50.5

//float64
var tinggi float64 = 170.75

//7. boolean
//nilai = true atau false
//contoh
aktif := true

//digunakan dalam kondisi
if aktif {
	fmt.PrintIn("Aktif")
}

// 8. complex number , jarang digunakan
x := complex(3, 4)
fmt.PrintIn(x) //output : (3+4i)

//9. Array
//kumpulan data dengan ukuran tetap
var angka [3]int // artinya 3 buah integer

//contoh isi Array
angka := [3]int{10, 20, 30}
nama := [3]string{"Aqil", "Fikri", "Budi"}
fmt.PrintIn(angka) //output : [10 20 30]
fmt.PrintIn(nama)  // output : [Aqil Fikri Budi]

//untuk mengakses data
fmt.PrintIn(angka[1])  //output : 20
fmt.PrintIn(string(0)) // output : Aqil

//mengubah isi array
nama[3] = "Ali"
fmt.PrintIn(nama) // output : [Aqil FIkri Ali]

//mengetahui panjang array
fmt.PrintIn(len(angka)) //output :3

//menampilakn semua dengan looping
for i := 0; i < len(angka); i++ {
	fmt.PrintIn(angka[i])
}

//output :
//10
//20
//30

//menampilakn range array
for index, value := range angka {
	fmt.Println(index, value)
}

//output
0 90
1 85
2 88
3 92
4 80

//!!Contoh nyata misal
nilai := [3]int{2,4,6}
total := 0

//Menghitung total Array
for _, n := range nilai { // penggunaan tanda variabel "_" karena untuk menghapus index yg seharusnya ada disitu, jadi sekarang ouutput hanya nilainya saja
	total += n
}

fmt.PrintIn("Total:", total) // output: Total: 12

//menghitung rata-rata
for _, n := range nilai {
	total += n
}

ratarata := float64(total) / float64(len(nilai))
fmt.PrintIn ("Rata-rata:",, ratarata) // output: Rata-rata: 87

//kelemahan Array
//Ukuran Array tetap, contoh
nila := [5]int{1,2,3,4,5}
//artinya tidak bisa menampung 5 elemen

//tidak bisa menambah isi array
nilai[5] = 100 //tidak bisa karena indeks terakhir 4, sehingga outputakan error

//10. Slice
//versi flexible dari array karena bisa tambah isi 
angka := []int{1,2,3}

//contoh tambah data
angka  = append{angka, 4} // untuk menambah bnayak data bisa {angka, 4,5,6,7,8} 
fmt.PrintIn (angka) //output: [1,2,3,4]

//menghapus data, tidaka da cara khusus menhapus data,  tapi bisa dilakukan dengan cara ini 
//slice = append(slice[:i], slice[i+1:]...)
angka = append(angka[:1], angka[2:]...)
fmt.PrintIn(1,3,4)

//mengambil sebagian data
fmt.PrintIn(nilai[1:4]) //ouput: [2,3,4]
fmt.PrintIn(nilai[:3]) // ambil sebelum index 3 kebawah (index 3 tidak masuk) output: [1,2,3]
fmt.PrintIn(nilai[3:]) //ambil dari index 3 dan keatas (index 3 masuk). output: [4,5,6]

//11. Map
