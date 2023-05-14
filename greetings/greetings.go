// package greetings

// import "fmt"

// func Hello(name string) string {
// 	message := fmt.Sprintf("Hi, %v. Welcome", name)
// 	return message
// }

package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	// If no name is given return a n error.
	if name == "" {
		return "", errors.New("Empty Name!")
	}

	message := fmt.Sprintf("Hi, %v. Welcome", name)
	return message
}