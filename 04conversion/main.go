package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza App.")
	fmt.Println("Please rate our Pizza b/w 1 to 5.")

	reader := bufio.NewReader(os.Stdin)

	// keep reading input till ENTER.
	// using _ to store errors, as we don't care about it really.
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating: ", input)

	// let's try to do some manipulation on user's input.
	// Since, user input is of always type String, we convert input to int.
	// ParseFloat(input, bitsize)

	// numRating, err := strconv.ParseFloat(input, 64)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// Above code fails, along with `4` the `\n` is also being sent as innput.
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to rating!", numRating+1)
	}
}
