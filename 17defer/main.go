// Deferred functions are invoked just before the function returns i.e. at the end of function.
// In case of multiple deferred functions, LIFO is used.

package main

import "fmt"

func main() {
	defer fmt.Print("O\n")
	defer fmt.Print("F")
	defer fmt.Print("I")
	defer fmt.Print("L")
	myDefer()
	defer fmt.Println("\nDeferred works as: ")
	fmt.Println("\nHello to defer!")
	// myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
