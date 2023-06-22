package main

import "fmt"

func main() {
	fmt.Println("Learning maps in GoLang")

	languages := make(map[string]string)
	languages["js"] = "Javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("Ruby is short for: ", languages["rb"])

	delete(languages, "rb")
	fmt.Println("List of all languages: ", languages)

	// looping over maps
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	for _, value := range languages {
		// _ are ignored in golang
		fmt.Printf("value is %v\n", value)
	}
}
