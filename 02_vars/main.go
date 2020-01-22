package main

import "fmt"


func main(){
	// MAIN TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// using var
	// var name = "Jackson"
	var age int32 = 26
	const isCool = true

	// Shorthand
	// name := "Jackson"
	size := 1.3

	name, email := "Jackson","jackson@gmail.com"

	fmt.Println(name, age, isCool, size, email)
	fmt.Printf("%T\n", size)
}

