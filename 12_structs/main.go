package main

import (
	"fmt"
	"strconv"
)

// define a person struct
type Person struct {
	// firstName string
	// lastName string
	// city string
	// gender string
	// age int

	firstName, lastName, city, gender string
	age int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// // Init person using struct
	person1 := Person{firstName: "Patrik", lastName: "Schorn", city: "Dortmund", gender: "m", age: 26}
	person2 := Person{firstName: "Stephanie", lastName: "Miller", city: "Dortmund", gender: "f", age: 18}
	// // Alternative
	// // person2 := Person{"Patrik", "Schorn", "Dortmund", "m", 26}

	// fmt.Println(person1)
	// // fmt.Println(person2)
	// fmt.Println(person1.city)
	// person1.age++
	// fmt.Println(person1)

	fmt.Println(person2.greet())
	person2.hasBirthday()
	
	person2.getMarried(person1.lastName)
	person1.getMarried("Laurens")
	fmt.Println(person2.greet())
	fmt.Println(person1.greet())
}
