package main

import "fmt"

func main() {
	//Init statement, dieksekusi sebelum perulangan pertama kali dilakukan.
	//     | Condition expression, melakukan pengecekan sebelum melakukan setiap perulangan.
	//     |      |     Post statment, dieksekusi diakhir setiap perulangan.
	//     |      |      |
	//     v      v      v
	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}

	y := 1
	for y < 10 {
		fmt.Println(y)
		y++
	}
}
