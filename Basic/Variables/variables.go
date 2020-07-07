package main

import "fmt"

// Variable didefinisikan dengan statment `var`.
// Variable bisa di dalam package atau di dalam fungsi.
var nama string
var umur int

func main() {
	umur = 21
	nama = "xan"
	var hobby = "makan"
	// Inisialisasi variable bisa dilakukan di satu statemen var.
	// Inisialisasi variable tidak perlu ditulis type datanya.
	var ukuranBaju, tahunLahir = "s", 14
	fmt.Println(nama, "umur", umur, hobby)
	fmt.Println(ukuranBaju, tahunLahir)
	//Operator `:=` digunakan untuk memperpendek cara mendeklarasikan variable.
	//Operator `:=` hanya bekerja didalam fungsi.
	umur := 23
	fmt.Println(umur)
	nama := "doctr"
	test(nama)
}

func test(namabaru string) {
	fmt.Println(nama, umur)
	fmt.Println(namabaru)
}
