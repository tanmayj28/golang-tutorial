package main

import "fmt"

type Rectangle struct {
	Length float64
	Width  float64
}

// YOUR CODE HERE:
// Define a Rectangle struct type with Length and Width
// fields, each of which has a type of float64.

func (rect Rectangle) Area() float64 {
	return rect.Length * rect.Width
}

// YOUR CODE HERE:
// Define an Area method on the Rectangle type. It should
// accept no parameters (other than the receiver parameter).
// It should return a float64 value calculated by multiplying
// the receiver's Length by its Width.

func (rect Rectangle) Perimeter() float64 {
	return rect.Length*2 + 2*rect.Width
}

// YOUR CODE HERE:
// Define a Perimeter method on the Rectangle type. It should
// accept no parameters. It should return a float64 value
// representing the receiver's perimeter (2 times its Length
// plus 2 times its Width).

func main() {
	// Once you've defined the above code correctly,
	// this code should compile and run.
	var myRectangle Rectangle
	myRectangle.Length = 2
	myRectangle.Width = 3
	fmt.Println("Area:", myRectangle.Area())           // => Area: 6
	fmt.Println("Perimeter:", myRectangle.Perimeter()) // => Perimeter: 10
}
