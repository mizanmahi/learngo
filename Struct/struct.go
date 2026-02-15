package main

import "fmt"

type user struct {
	name  string
	email string
}

func main() {

	// john := user{name: "john", email: "john@mail.com"}

	// john.email = "mizan@mail.com"

	// fmt.Println(john.email)

	var user1 user

	user1.name = "john"
	user1.email = "test@mail.com"

	fmt.Println(user1)

}

// crating struct type
// Different ways to create struct
// using printf to print struct (%+v)
// updating struct fields
// Accessing struct fields

// embedding struct
// receiver function
