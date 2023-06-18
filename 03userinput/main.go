package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	welcome := "Welcome to userimput program"
	fmt.Println(welcome)

	// reader is listening and reading here
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of Tottenham:")

	// comma ok || err ok
	// use _ for outputs you don't care about

	// read till someone hits \n
	// input, err := reader.ReadString('\n')
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)

	// all input seem to be string, this could be issue if we want numbers, boolean etc.
	fmt.Printf("Type of rating is %T", input)
}
