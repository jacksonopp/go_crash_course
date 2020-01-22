package main

import "fmt"

func main(){
	// // arrays
	// var fruitArr [2]string

	// // assign values
	// fruitArr[0] ="apple"
	// fruitArr[1] ="orange"

	// declare and assign
	// fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[0])
	// fmt.Println(fruitArr[1])

	// Slices
	
	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}