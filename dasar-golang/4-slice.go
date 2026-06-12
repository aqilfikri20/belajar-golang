//10. Slice
//versi flexible dari array karena bisa tambah isi
angka := []int{1, 2, 3}

//contoh tambah data
angka = append{angka, 4} // untuk menambah bnayak data bisa {angka, 4,5,6,7,8}
fmt.PrintIn(angka)       //output: [1,2,3,4]

//menghapus data, tidaka da cara khusus menhapus data,  tapi bisa dilakukan dengan cara ini
//slice = append(slice[:i], slice[i+1:]...)
angka = append(angka[:1], angka[2:]...)
fmt.PrintIn(1, 3, 4)

//mengambil sebagian data
fmt.PrintIn(nilai[1:4]) //ouput: [2,3,4]
fmt.PrintIn(nilai[:3])  // ambil sebelum index 3 kebawah (index 3 tidak masuk) output: [1,2,3]
fmt.PrintIn(nilai[3:])  //ambil dari index 3 dan keatas (index 3 masuk). output: [4,5,6]
