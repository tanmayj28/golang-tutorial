package main

import "fmt"

func main() {
	fmt.Println("Loops in Go!")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	for day := 0; day < len(days); day++ {
		fmt.Println("Day is: ", days[day])
	}

	for day := range days {
		fmt.Println("Day is:: ", days[day])
	}

	for index, day := range days {
		fmt.Printf("Index is %v and Day is %v\n", index, day)
	}

	someValue := 1
	for someValue < 10 {
		fmt.Println("Some value is:", someValue)
		if someValue == 5 {
			break
		}
		someValue++
	}

	goToValue := 1
	for goToValue < 10 {
		fmt.Println("Some value is:", goToValue)
		if goToValue == 3 {
			goto lco
			// auto breaks the loop chain
		}
		goToValue++
	}

lco:
	fmt.Println("This is Go To LCO!")
}
