package main

import "fmt"

func main() {
	var input int
	fmt.Scanln(&input)
	// If dapat diawali dengan short statement, yang akan dieksekusi sebelum kondisi.
	// Variable yang dideklarasikan di statement hanya berlaku didalam blok if.
	if x := input; x == 1 {
		fmt.Println("njay")
	} else {
		// Variable yang dideklarasi dalam short statement berlaku juga dalam blok else.
		fmt.Println(x, "salah")
	}
}
