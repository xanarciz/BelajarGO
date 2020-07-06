// Semua program GO ditulis dalam sebuah package.
package main

// Import bisa dilakukan untuk menggunakan package lain.
import "fmt"

// Program pertama kali berjalan dalam package main.
func main() {
	// Nama yang diekspor diawali oleh kapital. Misal package fmt menggunakan fungsi println, maka ditulis
	// fmt.Println
	fmt.Println("test")
}
