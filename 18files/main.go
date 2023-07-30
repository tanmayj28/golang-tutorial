package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golang!")
	content := "This needs to be stored in a file."

	file, err := os.Create("./myTextFile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	// returns lenght if file gets created
	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length of created file is: ", length)
	defer file.Close()
	readFile("./myTextFile.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Text data inside the file is: \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		// panic() shuts down program ans shows the error
		panic(err)
	}
}
