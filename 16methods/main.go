// functions defined inside of Structs (since we dont have classes in golang) are called methods

package main

import "fmt"

// first letter capital means, they are accessible outside
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Welcome to methods.")
	fmt.Println("Getting started with Structs")
	user_tj := User{"Tanmay", "tanmaj28@gmail.com", true, 27}
	// just prints the values and not the keys
	fmt.Println(user_tj)

	// prints in key:value style, very commonly used.
	fmt.Printf("User details are: %+v\n", user_tj)
	fmt.Printf("User name is: %v\n", user_tj.Name)
	fmt.Printf("User age is: %v\n", user_tj.Age)

	user_tj.GetStatus()
	user_tj.RemoveMail() // doesn't actually update user_tj
	fmt.Printf("User details are: %+v\n", user_tj)
}

func (u User) GetStatus() {
	fmt.Println("Is user active? ", u.Status)
}

// updates the copy of the user object, send reference to actually update the user.
func (u User) RemoveMail() {
	u.Email = ""
	fmt.Println("User updated!", u)
}
