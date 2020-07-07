package main

import "fmt"

func main() {
	fmt.Println(tambah(12, 12, "test"))
	//Fungsi bisa menghasilkan beberapa hasil return.
	buah1, buah2 := buah("jambu", "pisang")
	fmt.Println(buah1)
	fmt.Println(buah2)
	fmt.Println(nakedreturn(2))
	nakedint, nakedstring := nakedreturn(2)
	fmt.Println(nakedint)
	fmt.Println(nakedstring)
}

// Fungsi bisa memiliki 0 argumen atau lebih
// a dan b memiliki type data integer
func tambah(a, b int, c string) int {
	return a + b
}

func buah(d, e string) (string, string) {
	return d, e
}

//naked return adalah adalah argumen return yang tidak memiliki argumen
func nakedreturn(x int) (y int, j string) {
	y = x + 20
	j = "andi"
	return
}
