package main

import "fmt"

func main() {
	fmt.Println("test")
	defer fmt.Println("wkwkw")
	fmt.Println("test 2")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
