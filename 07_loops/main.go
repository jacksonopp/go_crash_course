package main

import "fmt"

func main() {
	// long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		// i = i + 1
		i++
	}

	// short method

	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}
}
