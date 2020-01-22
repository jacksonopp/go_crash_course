package main

import "fmt"

func main() {
	// define maps
	// emails := make(map[string]string)

	// // assign key values
	// emails["Bob"] = "bob@gmail.com"
	// emails["sharon"] = "sharon@gmail.com"

	// declare map and add kv
	emails := map[string]string{"bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}
	emails["Mike"] = "Mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// Delete from map
	delete(emails, "bob")
	fmt.Println(emails)
}
