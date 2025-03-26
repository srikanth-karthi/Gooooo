package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")

	f := "apple"
	fmt.Printf("Type of f: %T\n", f) // Correct way to print the type of a variable

	var e bool
	fmt.Println(e)
	fmt.Println(f)
}
