package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices in golang")

	var fruitList = []string{"Apple", "Orange", "Peach"}
	fmt.Printf("Type of fruitList is: %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")

	fmt.Println("FruitList is: ", fruitList)

	fruitList = append(fruitList[1:]) // range in goLang, starts array from index = 1
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3]) // range in goLang, starts array from index = 1 till index = 2
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3]) // range in goLang, starts array from index = 0 till index = 2
	fmt.Println(fruitList)

	highScores := make([]int, 4) // init with (type, size)
	highScores[0] = 90
	highScores[1] = 89
	highScores[2] = 99
	highScores[3] = 60
	// highScores[4] = 60 // This will fail as we have already specified the size.
	highScores = append(highScores, 12, 10) // This won't fail, as append reallocates the memory

	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores) // sorts array, available in slices not in array
	fmt.Println(sort.IntsAreSorted(highScores))
	fmt.Println(highScores)

	// ------------------------how to remove value from slices based on index------------------------

	var courses = []string{"reactjs", "golang", "ruby", "python", "swift"}
	fmt.Println(courses)

	var index = 3
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
