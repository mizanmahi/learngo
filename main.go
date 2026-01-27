package main

import "fmt"

// isAdmin := true // wrong

func main() {

	var name string = "next level"
	rating := 4.7899434
	formattedString := fmt.Sprintf("My name is %s and my rating is %.2f ", name, rating)
	fmt.Println(formattedString)

}
