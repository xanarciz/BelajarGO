package main

import "fmt"

func main() {
	var input string
	fmt.Scan(&input)
	switch x := input; x {
	case "test":
		fmt.Println("njir test")
	case "dev":
		fmt.Println("ok dev")
	default:
		fmt.Println(">;(")
	}
}
