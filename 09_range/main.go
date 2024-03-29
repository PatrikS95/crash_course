package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 2}

	// loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// not using Index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum: ", sum)

	// Range with map
	emails := map[string]string{"Bob":"bob@gmail.de", "Sharon": "sharon@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
	for k := range emails {
		fmt.Printf("%s\n", k)
	}
	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
