package main

import "fmt"

func tambah(a int, b int) int {
    return a + b
}

func main() {
    hasil := tambah(5, 7)
    fmt.Println("Hasil penjumlahan:", hasil)
}
