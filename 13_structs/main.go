package main

import (
	"fmt"
	"strconv"
)

// define person struct
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

// greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getsMarried (pointer reciever)
func (p *Person) getsMarried(newLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = newLastName
	}
}

func main() {
	// init person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Chicago", gender: "f", age: 20}

	person2 := Person("Bob", "Johnson", "Boston", "m", 30)

	// person1 := Person{"Samantha", "Smith", "Chicago", "f", 20}
	// fmt.Println(person1)

	// get single field
	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1)

	fmt.Println(person1.greet())
	person1.hasBirthday()
	person1.getsMarried("Williams")
	fmt.Println(person1.greet())

	person2.getsMarried("Thomson")
	fmt.Println(person2.greet())

}
