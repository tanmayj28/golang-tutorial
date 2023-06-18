package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers!")

	// If no value assigned to Pointer it will be <nil>.
	var ptr *int
	fmt.Println("Value of Pointer is: ", ptr)

	// A reference means &, i.e. it will pass the reference of the actual variable
	myNumber := 23
	var ptrMyNumber = &myNumber
	fmt.Println("Value of myNumber Pointer is: ", ptrMyNumber)
	fmt.Println("Value of myNumber using Pointer is: ", *ptrMyNumber)

	// This shows how passing the address of value makes sure the operations actually happen on the values and not on the copies of the values
	*ptrMyNumber = *ptrMyNumber * 2
	fmt.Println("New value is: ", myNumber)
}
