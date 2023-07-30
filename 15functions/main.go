// functions defined outside of scopes are called functions.
// can't define functions inside functions.
// sequence doesn't matter
// anonymous functions exist.

package main

import "fmt"

func adder(inta int, intb int) int {
	return inta + intb
}

func proAdder(ints ...int) (int, string) {
	total := 0
	for _, val := range ints {
		total += val
	}
	return total, "Pro Adder returns:"
}

func main() {
	fmt.Println("Welcome to functions in Golang!")
	greeter()
}

func greeter() {
	fmt.Println("Namaste from Golang!")
	result := adder(3, 4)
	fmt.Println("adder returns:", result)

	proRes, proText := proAdder(3, 4, 6)
	fmt.Println(proText, proRes)
}

// func () {
// 	fmt.Println("Anonymous function!")
// }();
