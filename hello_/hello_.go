package main
import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// message, err := greetings.Hello("")
	// message, err := greetings.Hello("Tanmay")

	// A slice of names.
	names := [] string {"Kane", "Son", "Kulu"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(message)
	fmt.Println(messages)
}
