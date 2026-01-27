package main

import "fmt"

// func makeCoffee() {

// }

func main() {
	// anonymous function
	// coffeeBanao := func() {
	// 	fmt.Printf("Making coffee")
	// }

	// coffeeBanao()

	// IIFE
	func(coffeeType string) {
		fmt.Printf("Making hot %s....", coffeeType)
	}("Latte")
}
