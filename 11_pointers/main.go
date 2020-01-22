package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// read from pointer: use * to read val

	fmt.Println(*b)

	// change val with pointer
	*b = 10
	fmt.Println(a)
}
