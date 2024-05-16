package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to web requests in Golang!")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	// callers responsibility to close the response!
	defer response.Body.Close()

	fmt.Printf("Success response! It is of type %T\n", response)

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Databyte is:", string(databytes))
}
