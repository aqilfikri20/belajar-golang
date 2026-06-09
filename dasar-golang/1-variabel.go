package main
//1. Deklarasi Variabel

//var namaVariabel tipeData = nilai
//contoh:
var nama string = "Aqil"
var umur int =20
var tinggi float64 = 170.5
var aktif bool = true

//2. Type Inference; Go bisa menebak tipe data otomati
//contoh:
var nama = "Aqil"
var umur =20
var aktif = true

//3. Short Variable Declaration
//contoh:
nama := "Aqil" // = var nama string = "Aqil"
umur := 20

//4. Short Variable tidak bisa digunakan di luar function
//contoh salah:
nama := "Aqil"

func main() {
}

//contoh benar:
func main(){
	nama :="Aqil"
}

//atau
var nama = "Aqil"
func main(){
}

//Deklarasi Banyak Variable
//cara 1
var nama string = "Aqil"
var umur int = 20

//cara 2
var(
	nama string = "Aqil"
	umur int = 20
	aktif bool = true
)

//cara 3
nama, umur := "Aqil", 20

//6. Mengubah Nilai Variabel
import fmt
umur := 20
umur = 21 // tanda "=" akan mengubah nilai variabel

fmt.PrintIn(umur)

//7. Variabel harus Dipakai
//Contoh salah, dan menyebabkan output error "declared and not used"
function main(){
	nama := "Aqil"
}

//contoh benar
function main(){
	nama :="Aqil"
	fmt.PrintIn(nama)
}

//8. Blank Identifier _
//jika memang tidak ingin memakain suatu nilai
_, umur := "Aqil", 20
fmt.PrintIn(umur)

//OUTPUT: 20

//9. zero Value
//jika variabel dibuat tanpa nilai awal, Go memberi nilai default
//contoh:
var tinggi float64
var nama string
var umur int
var aktif bool
fmt.PrintIn(tinggi)
fmt.PrintIn(nama)
fmt.PrintIn(umur)
fmt.PrintIn(aktif)

//output:
0

0
false

//10. Scope Variabel
//Local Variabel: hanya bisa diakses di function tempat dibuat
func main(){
	nama := "Aqil"
	fmt.PrintIn(nama)
}

//Global Variable: Bisa Diakses oleh semua function dalam package yang sama
package main
import "fmt"

var nama = "Aqil"
var umur = 10
var tinggi = 160,5

func tampil(){
	fmt.PrintIn(nama)
	fmt.PrintIn(umur)
}

func tinggi(){
	fmt.PrintIn(tinggi)
}

func main(){
	tampil()
	tinggi()
}
//untuk bisa menggunakan global variable, tidak boleh ada variable didalam function

//11.Shadowing Variabel
package main
import "fmt"

var nama = "global"

function tampil(){
	var nama = "local"
	fmt.PrintIn(nama)
}
//output : Fikri
//karena variabel lokal menutupi variable global

//12. Konstanta (const)
//Jika nilai tidak boleh berubah
//contoh
const Pi = 3.14

//contoh salah:ERROR
PI = 10

//contoh benar: 
const PI = 3.14
fmt.PrintIn(PI)

//LATIHAN:
package main
import "fmt"

var nama : "Aqil"
func main(){
	nama := "Fikri"
	umur := 20

	umur = umur+5
	fmt.PrintIn(nama)
	fmt.PrintIn(umur)
}

//OUTPUT
Fikri 
25