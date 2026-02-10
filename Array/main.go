package main

import "fmt"

func main() {

	var numbers [6]int

	numbers[0] = 42
	numbers[1] = 7

	numbers[5] = 100

	// fmt.Println(numbers)

	fmt.Println("Length of numbers is:", len(numbers))

}
