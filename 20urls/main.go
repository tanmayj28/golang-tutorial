package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs"

func main() {
	fmt.Println("Welcome to handling URLs in GoLang!")
	result, _ := url.Parse(myUrl)

	fmt.Println("URL scheme is: ", result.Scheme)
	fmt.Println("URL scheme is: ", result.Host)
	fmt.Println("URL scheme is: ", result.Path)
	fmt.Println("URL scheme is: ", result.RawQuery)
	fmt.Println("URL scheme is: ", result.Port())

	qParams := result.Query()

	fmt.Println("Coursename from query params is: ", qParams["coursename"])

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "google.com",
		Path:    "/images",
		RawPath: "user=tanmay",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println("New URL formed is: ", anotherUrl)
}
