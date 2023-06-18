package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time Study of golang!")
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(1995, time.November, 21, 5, 0, 0, 0, time.UTC)
	fmt.Println(createDate.Format("01-02-2006 Monday"))

	// For random uuid values:
	fmt.Println(presentTime.UnixNano())
}
