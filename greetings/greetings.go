// package greetings

// import "fmt"

// func Hello(name string) string {
// 	message := fmt.Sprintf("Hi, %v. Welcome", name)
// 	return message
// }

// -----------------------------------------

// package greetings

// import (
// 	"errors"
// 	"fmt"
// )

// func Hello(name string) (string, error) {
// 	// If no name is given return a n error.
// 	if name == "" {
// 		return "", errors.New("Empty Name!")
// 	}

// 	message := fmt.Sprintf("Hi, %v. Welcome", name)
// 	return message, nil
// }

// -----------------------------------------

package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return name, errors.New("Empty Name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := [] string {
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	// Return a randomly selected message format by specifying a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}

// -------------------------------------------------------------------------------
// Hellos returns a map that associates each of the named people with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)

	// Loop through the received slice of names, calling the Hello function to get a message for each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with the name.
		messages[name] = message
	}

	return messages, nil
}
