// package main

// import "fmt"

// func deferred(result int) {
// 	fmt.Println("deferred result:", result)
// }

// func example() int {
// 	result := 10

// 	defer deferred(result) // deferred(10)

// 	fmt.Println("I am from example fn:", result)

// 	// result += 100
// 	result = result + 100

// 	return result
// }

// func main() {
// 	example()
// }

package main

import "fmt"

func deferred(result int) {
	fmt.Println("deferred result:", result)
}

func example() int {
	result := 10

	defer deferred(result) // deferred(10)

	fmt.Println("I am from example fn:", result)

	// result += 100
	result = result + 100

	return result
}

func main() {
	example()
}
