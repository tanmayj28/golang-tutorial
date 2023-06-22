package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in GoLang!")
	var fruitList [4]string // important to specify the size here

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	// fruitList[2] = "Peach"
	fruitList[3] = "Orange"

	fmt.Println("Fruit List is: ", fruitList)                // has a space at [2] index
	fmt.Println("Length of Fruit List is: ", len(fruitList)) // gives 4 even if 3 elements inserted

	var vegList = [3]string{"Potato", "Beans", "Onion"}
	fmt.Println("Veg List is: ", vegList)
	fmt.Println("Length of Veg List is: ", len(vegList))
}
