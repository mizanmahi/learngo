package main

// isAdmin := true // wrong

func main() {

	user_name := "alice" // snake case, correct syntax
	isAdmin := true      // camel case, correct syntax, preferred in Go

	// invalid variable names
	// 1st-name := "bob" // hyphen not allowed
	// 2ndName := "charlie" // cannot start with a number
	// var float := 3.14 // 'float' is a reserved keyword

	println("User Name:", user_name)
	println("Is Admin:", isAdmin)

}
