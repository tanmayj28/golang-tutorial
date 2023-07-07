package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in Go!")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spots!")
	case 3:
		fmt.Println("You can move 3 spots!")
		fallthrough // There are breaks ON by default in Go. If you want to run case: 4, fallthrough will help. If you want further values to be printed as well just add fallthrough's down below too.
	case 4:
		fmt.Println("You can move 4 spots!")
	case 5:
		fmt.Println("You can move 5 spots!")
	case 6:
		fmt.Println("You can move 6 spots and roll dice again!")
	}
}
