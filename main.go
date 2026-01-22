package main

import "fmt"

// isAdmin := true // wrong

func main() {

	// name := "Mizan" // short way, most used in real projects
	// name string := "Mizan" // wrong way

	// var name string = "Mizan"
	// var name = "mizan"

	// grouped variables
	var (
		name string = "Mizan"
		age  int    = 30
	)

	// multiple variables declaration
	// var x, y int

	// x = 25
	// y = 30
	var x, y string = "mizan", "next level"

	const point int = 5

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(x, y, point)

}
