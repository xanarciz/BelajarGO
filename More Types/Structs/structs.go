package main

import "fmt"

//Struct adalah sekumpulan fields.
type member struct {
	nama  string
	noKtp int
}

func main() {
	andi := member{"andi", 123}
	//Field dappat di akses menggunakan tanda titik.
	fmt.Println(andi.nama)
	fmt.Println(andi.noKtp)
}
