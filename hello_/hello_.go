package main
import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("")
	message := greetings.Hello("Tanmay")
	fmt.Println(message)
}