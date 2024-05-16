// Fill in the blanks so this program compiles and produces
// the output shown.
package main

import "fmt"

func main() {
	// Create a variable to hold a slice of ints.
	var primes []int
	// Create a slice with 2 elements.
	primes = make([]int, 2)
	// Assign values to the first 2 elements.
	primes[0] = 2
	primes[1] = 3
	// Add a third element to the end of the slice.
	primes = append(primes, 5)
	fmt.Println(primes) // => [2 3 5]

	// Write a map literal with int keys and string values.
	elements := map[int]string{1: "H", 2: "He", 3: "Li"}
	// Loop over each key/value pair in the map.
	for atomicNumber, symbol := range elements {
		fmt.Println(atomicNumber, symbol)
	}
	// => 1 H
	// => 2 He
	// => 3 Li
}
