package main

import "fmt"

func main() {
	fmt.Println("If else in GO!")
	voterAge := 18
	var result string

	if voterAge >= 18 {
		result = "Valid voter"
	} else {
		result = "Invalid voter"
	}
	fmt.Println(result)
}
