package main

import "fmt"

func main() {
	var angka [5]int
	angka[0] = 10
	angka[1] = 20
	angka[2] = 30

	fmt.Println(angka)
	fmt.Println(angka[1])

	nama := [3]string{"Aqil", "Budi", "Citra"}
	fmt.Println(nama)
}
