package main

import "fmt"

func main() {
	i := 2
	switch i {
	case 0:
		fmt.Println("Zero")
		fmt.Println("Zero")
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	default:
		fmt.Println("Unknown")
	}
}
