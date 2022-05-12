package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int	int8	int16	int32	int64
	// uint	uint8	uint16	uint32	uint64
	// byte - alias for uint8
	// rune - alias for int32
	// float32	float64
	// complex64	complex128

	// Using var
	// var name string = "Patrik"
	// Alternative:
	name, email2 := "Patrik", "patrik.schorn@tu-dortmund.de"
	alter := 26
	const isCool = true
	größe := 184.5
	email := "patrik.schorn@tu-dortmund.de"

	fmt.Println(name, alter, isCool, größe, email, email2)
	fmt.Printf("%T\n", größe)
}
