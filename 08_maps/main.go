package main

import "fmt"

func main() {
	// // define map
	// emails := make(map[string]string)

	// // assign key values
	// emails["Bob"] = "bob@gmail.com"
	// emails["Peter"] = "peter@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	// Declare map and add key:values
	emails := map[string]string{"Bob":"bob@gmail.de", "Sharon": "sharon@gmail.com"}

	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)

}
