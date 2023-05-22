package main

import "fmt"

const PublicVariable = 420 // since it starts with capital it is PUBLIC

func main() {
	var username string = "tanmay"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallValue uint8 = 255 // can only store 8 bit integers i.e. <= 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T \n", smallValue)

	var smallFloat float32 = 2.123456789098765 // set of 32 bit digits
	var bigFloat float64 = 2.123456789098765   // set of 64 bit digits
	fmt.Println(smallFloat)
	fmt.Println(bigFloat)
	fmt.Printf("Variable is of type: %T and %T \n", smallFloat, bigFloat)

	// Default values and aliases
	var defaultValue int
	fmt.Println(defaultValue)
	fmt.Printf("Variable is of type: %T \n", defaultValue)

	// impicit types
	var website = "google.com"
	fmt.Println(website) // Lexer does the type setting for us here but it won't allow you to update it's type later.

	// No var style using Walrus Operator
	// This is only allowed inside methods and not in globa; scopes etc.
	numberOfUsers := 200
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type: %T \n", numberOfUsers)

	// Constans
	fmt.Println(PublicVariable)
	fmt.Printf("Variable is of type: %T \n", PublicVariable)
}
